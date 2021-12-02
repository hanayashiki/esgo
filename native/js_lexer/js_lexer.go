package js_lexer

import (
	"fmt"
	"unicode/utf8"

	"github.com/hanayashiki/esgo/native/logger"
)

type Lexer struct {
	source                logger.Source
	currentCodeOffset     int
	nextCodeOffset        int
	currentCode           rune
	currentToken          T
	currentIdentifierName string
	currentLocation       logger.Location
}

func NewLexer(source logger.Source) Lexer {
	lexer := Lexer{
		source: source,
		currentLocation: logger.Location{
			Offset: 0,
			Row: 1,
			Col: 0,
		},
	}

	lexer.step()
	lexer.Next()

	return lexer
}

func (lexer *Lexer) step() {
	if lexer.nextCodeOffset >= len(lexer.source.Contents) {
		lexer.currentCode = -1
		return
	}

	code, size := utf8.DecodeRuneInString(lexer.source.Contents[lexer.nextCodeOffset:])

	if code == utf8.RuneError {
		panic("Got invalid utf-8 input")
	}

	if size == 0 {
		lexer.currentCode = -1
	} else {
		lexer.currentCode = code
		lexer.currentCodeOffset = lexer.nextCodeOffset
		lexer.nextCodeOffset += size

		if code == '\n' {
			lexer.currentLocation.Row++
			lexer.currentLocation.Col = 1
		} else {
			lexer.currentLocation.Col += 1
		}
		lexer.currentLocation.Offset = lexer.currentCodeOffset
	}
}

func (lexer *Lexer) PeekToken() T {
	return lexer.currentToken
}

func (lexer *Lexer) Next() {
	for {
		switch lexer.currentCode {
		case -1:
			lexer.currentToken = TEndOfFile
			return
		case '(':
			lexer.step()
			lexer.currentToken = TOpenParen
			return
		case ')':
			lexer.step()
			lexer.currentToken = TCloseParen
			return
		case '[':
			lexer.step()
			lexer.currentToken = TOpenBracket
			return
		case ']':
			lexer.step()
			lexer.currentToken = TCloseBracket
			return
		case '{':
			lexer.step()
			lexer.currentToken = TOpenBrace
			return
		case '}':
			lexer.step()
			lexer.currentToken = TCloseBrace
			return
		case ',':
			lexer.step()
			lexer.currentToken = TComma
			return
		case ':':
			lexer.step()
			lexer.currentToken = TColon
			return
		case ';':
			lexer.step()
			lexer.currentToken = TSemicolon
			return
		case '@':
			lexer.step()
			lexer.currentToken = TAt
			return
		case '~':
			lexer.step()
			lexer.currentToken = TTilde
			return
		case '?':
			lexer.nextQuestionToken()
			return
		case '%':
			lexer.step()
			if lexer.currentCode == '=' {
				lexer.step()
				lexer.currentToken = TPercentEquals
			} else {
				lexer.currentToken = TPercent
			}
			return
		case '^':
			lexer.step()
			if lexer.currentCode == '=' {
				lexer.step()
				lexer.currentToken = TCaretEquals
			} else {
				lexer.currentToken = TCaret
			}
			return
		case '&':
			lexer.nextAmpersandToken()
			return
		case '|':
			lexer.nextBarToken()
			return
		case '+':
			lexer.nextPlusToken()
			return
		case '-':
			lexer.nextMinusToken()
			return
		default:
			if isWhitespace(lexer.currentCode) {
				lexer.step()
				continue
			}

			if isIdentifierStart(lexer.currentCode) {
				lexer.currentToken = TIdentifier
				start := lexer.currentCodeOffset
				lexer.step()

				for isIdentifierPart(lexer.currentCode) {
					lexer.step()
				}
				name := lexer.source.Contents[start:lexer.nextCodeOffset]

				keyword, exists := Keywords[name]

				if exists {
					lexer.currentToken = keyword
				} else {
					lexer.currentToken = TIdentifier
					lexer.currentIdentifierName = name
				}

				return
			}

			panic(fmt.Sprintf("Unrecognized code point %v\n", lexer.currentCode))
		}
	}
}

func (lexer *Lexer) nextQuestionToken() {
	lexer.step()
	if lexer.currentCode == '?' {
		lexer.step()
		if lexer.currentCode == '=' {
			lexer.step()
			lexer.currentToken = TQuestionQuestionEquals
		} else {
			lexer.currentToken = TQuestionQuestion
		}
	} else if lexer.currentCode == '.' {
		lexer.step()
		lexer.currentToken = TQuestionDot
	} else {
		lexer.currentToken = TQuestion
	}
}

func (lexer *Lexer) nextAmpersandToken() {
	lexer.step()
	if lexer.currentCode == '&' {
		lexer.step()
		if lexer.currentCode == '=' {
			lexer.step()
			lexer.currentToken = TAmpersandAmpersandEquals
		} else {
			lexer.currentToken = TAmpersandAmpersand
		}
	} else {
		lexer.currentToken = TAmpersand
	}
}

func (lexer *Lexer) nextBarToken() {
	lexer.step()
	if lexer.currentCode == '|' {
		lexer.step()
		if lexer.currentCode == '=' {
			lexer.step()
			lexer.currentToken = TBarBarEquals
		} else {
			lexer.currentToken = TBarBar
		}
	} else {
		lexer.currentToken = TBar
	}
}

func (lexer *Lexer) nextPlusToken() {
	// +, +=, ++
	lexer.step()
	if lexer.currentCode == '=' {
		lexer.step()
		lexer.currentToken = TPlusEquals
	} else if lexer.currentCode == '+' {
		lexer.step()
		lexer.currentToken = TPlusPlus
	} else {
		lexer.currentToken = TPlus
	}
}

func (lexer *Lexer) nextMinusToken() {
	// -, -=, --
}

func (lexer *Lexer) C() string {
	return lexer.currentIdentifierName
}
