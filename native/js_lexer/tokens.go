// Gratefully copied from github.com/evanw/esbuild/ 
//
// MIT License

// Copyright (c) 2020 Evan Wallace

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.


package js_lexer

type T uint

// If you add a new token, remember to add it to "tokenToString" too
const (
	TEndOfFile T = iota
	TSyntaxError

	// "#!/usr/bin/env node"
	THashbang

	// Literals
	TNoSubstitutionTemplateLiteral // Contents are in lexer.StringLiteral ([]uint16)
	TNumericLiteral                // Contents are in lexer.Number (float64)
	TStringLiteral                 // Contents are in lexer.StringLiteral ([]uint16)
	TBigIntegerLiteral             // Contents are in lexer.Identifier (string)

	// Pseudo-literals
	TTemplateHead   // Contents are in lexer.StringLiteral ([]uint16)
	TTemplateMiddle // Contents are in lexer.StringLiteral ([]uint16)
	TTemplateTail   // Contents are in lexer.StringLiteral ([]uint16)

	// Punctuation
	TAmpersand
	TAmpersandAmpersand
	TAsterisk
	TAsteriskAsterisk
	TAt
	TBar
	TBarBar
	TCaret
	TCloseBrace
	TCloseBracket
	TCloseParen
	TColon
	TComma
	TDot
	TDotDotDot
	TEqualsEquals
	TEqualsEqualsEquals
	TEqualsGreaterThan
	TExclamation
	TExclamationEquals
	TExclamationEqualsEquals
	TGreaterThan
	TGreaterThanEquals
	TGreaterThanGreaterThan
	TGreaterThanGreaterThanGreaterThan
	TLessThan
	TLessThanEquals
	TLessThanLessThan
	TMinus
	TMinusMinus
	TOpenBrace
	TOpenBracket
	TOpenParen
	TPercent
	TPlus
	TPlusPlus
	TQuestion
	TQuestionDot
	TQuestionQuestion
	TSemicolon
	TSlash
	TTilde

	// Assignments (keep in sync with IsAssign() below)
	TAmpersandAmpersandEquals
	TAmpersandEquals
	TAsteriskAsteriskEquals
	TAsteriskEquals
	TBarBarEquals
	TBarEquals
	TCaretEquals
	TEquals
	TGreaterThanGreaterThanEquals
	TGreaterThanGreaterThanGreaterThanEquals
	TLessThanLessThanEquals
	TMinusEquals
	TPercentEquals
	TPlusEquals
	TQuestionQuestionEquals
	TSlashEquals

	// Class-private fields and methods
	TPrivateIdentifier

	// Identifiers
	TIdentifier     // Contents are in lexer.Identifier (string)
	TEscapedKeyword // A keyword that has been escaped as an identifer

	// Reserved words
	TBreak
	TCase
	TCatch
	TClass
	TConst
	TContinue
	TDebugger
	TDefault
	TDelete
	TDo
	TElse
	TEnum
	TExport
	TExtends
	TFalse
	TFinally
	TFor
	TFunction
	TIf
	TImport
	TIn
	TInstanceof
	TNew
	TNull
	TReturn
	TSuper
	TSwitch
	TThis
	TThrow
	TTrue
	TTry
	TTypeof
	TVar
	TVoid
	TWhile
	TWith
)

var tokenToString = map[T]string{
	TEndOfFile:   "end of file",
	TSyntaxError: "syntax error",
	THashbang:    "hashbang comment",

	// Literals
	TNoSubstitutionTemplateLiteral: "template literal",
	TNumericLiteral:                "number",
	TStringLiteral:                 "string",
	TBigIntegerLiteral:             "bigint",

	// Pseudo-literals
	TTemplateHead:   "template literal",
	TTemplateMiddle: "template literal",
	TTemplateTail:   "template literal",

	// Punctuation
	TAmpersand:                         "\"&\"",
	TAmpersandAmpersand:                "\"&&\"",
	TAsterisk:                          "\"*\"",
	TAsteriskAsterisk:                  "\"**\"",
	TAt:                                "\"@\"",
	TBar:                               "\"|\"",
	TBarBar:                            "\"||\"",
	TCaret:                             "\"^\"",
	TCloseBrace:                        "\"}\"",
	TCloseBracket:                      "\"]\"",
	TCloseParen:                        "\")\"",
	TColon:                             "\":\"",
	TComma:                             "\",\"",
	TDot:                               "\".\"",
	TDotDotDot:                         "\"...\"",
	TEqualsEquals:                      "\"==\"",
	TEqualsEqualsEquals:                "\"===\"",
	TEqualsGreaterThan:                 "\"=>\"",
	TExclamation:                       "\"!\"",
	TExclamationEquals:                 "\"!=\"",
	TExclamationEqualsEquals:           "\"!==\"",
	TGreaterThan:                       "\">\"",
	TGreaterThanEquals:                 "\">=\"",
	TGreaterThanGreaterThan:            "\">>\"",
	TGreaterThanGreaterThanGreaterThan: "\">>>\"",
	TLessThan:                          "\"<\"",
	TLessThanEquals:                    "\"<=\"",
	TLessThanLessThan:                  "\"<<\"",
	TMinus:                             "\"-\"",
	TMinusMinus:                        "\"--\"",
	TOpenBrace:                         "\"{\"",
	TOpenBracket:                       "\"[\"",
	TOpenParen:                         "\"(\"",
	TPercent:                           "\"%\"",
	TPlus:                              "\"+\"",
	TPlusPlus:                          "\"++\"",
	TQuestion:                          "\"?\"",
	TQuestionDot:                       "\"?.\"",
	TQuestionQuestion:                  "\"??\"",
	TSemicolon:                         "\";\"",
	TSlash:                             "\"/\"",
	TTilde:                             "\"~\"",

	// Assignments
	TAmpersandAmpersandEquals:                "\"&&=\"",
	TAmpersandEquals:                         "\"&=\"",
	TAsteriskAsteriskEquals:                  "\"**=\"",
	TAsteriskEquals:                          "\"*=\"",
	TBarBarEquals:                            "\"||=\"",
	TBarEquals:                               "\"|=\"",
	TCaretEquals:                             "\"^=\"",
	TEquals:                                  "\"=\"",
	TGreaterThanGreaterThanEquals:            "\">>=\"",
	TGreaterThanGreaterThanGreaterThanEquals: "\">>>=\"",
	TLessThanLessThanEquals:                  "\"<<=\"",
	TMinusEquals:                             "\"-=\"",
	TPercentEquals:                           "\"%=\"",
	TPlusEquals:                              "\"+=\"",
	TQuestionQuestionEquals:                  "\"??=\"",
	TSlashEquals:                             "\"/=\"",

	// Class-private fields and methods
	TPrivateIdentifier: "private identifier",

	// Identifiers
	TIdentifier:     "identifier",
	TEscapedKeyword: "escaped keyword",

	// Reserved words
	TBreak:      "\"break\"",
	TCase:       "\"case\"",
	TCatch:      "\"catch\"",
	TClass:      "\"class\"",
	TConst:      "\"const\"",
	TContinue:   "\"continue\"",
	TDebugger:   "\"debugger\"",
	TDefault:    "\"default\"",
	TDelete:     "\"delete\"",
	TDo:         "\"do\"",
	TElse:       "\"else\"",
	TEnum:       "\"enum\"",
	TExport:     "\"export\"",
	TExtends:    "\"extends\"",
	TFalse:      "\"false\"",
	TFinally:    "\"finally\"",
	TFor:        "\"for\"",
	TFunction:   "\"function\"",
	TIf:         "\"if\"",
	TImport:     "\"import\"",
	TIn:         "\"in\"",
	TInstanceof: "\"instanceof\"",
	TNew:        "\"new\"",
	TNull:       "\"null\"",
	TReturn:     "\"return\"",
	TSuper:      "\"super\"",
	TSwitch:     "\"switch\"",
	TThis:       "\"this\"",
	TThrow:      "\"throw\"",
	TTrue:       "\"true\"",
	TTry:        "\"try\"",
	TTypeof:     "\"typeof\"",
	TVar:        "\"var\"",
	TVoid:       "\"void\"",
	TWhile:      "\"while\"",
	TWith:       "\"with\"",
}

var Keywords = map[string]T{
	// Reserved words
	"break":      TBreak,
	"case":       TCase,
	"catch":      TCatch,
	"class":      TClass,
	"const":      TConst,
	"continue":   TContinue,
	"debugger":   TDebugger,
	"default":    TDefault,
	"delete":     TDelete,
	"do":         TDo,
	"else":       TElse,
	"enum":       TEnum,
	"export":     TExport,
	"extends":    TExtends,
	"false":      TFalse,
	"finally":    TFinally,
	"for":        TFor,
	"function":   TFunction,
	"if":         TIf,
	"import":     TImport,
	"in":         TIn,
	"instanceof": TInstanceof,
	"new":        TNew,
	"null":       TNull,
	"return":     TReturn,
	"super":      TSuper,
	"switch":     TSwitch,
	"this":       TThis,
	"throw":      TThrow,
	"true":       TTrue,
	"try":        TTry,
	"typeof":     TTypeof,
	"var":        TVar,
	"void":       TVoid,
	"while":      TWhile,
	"with":       TWith,
}

var StrictModeReservedWords = map[string]bool{
	"implements": true,
	"interface":  true,
	"let":        true,
	"package":    true,
	"private":    true,
	"protected":  true,
	"public":     true,
	"static":     true,
	"yield":      true,
}
