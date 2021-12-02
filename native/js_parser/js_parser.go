package js_parser

import (
	"github.com/hanayashiki/esgo/native/js_lexer"
	"github.com/hanayashiki/esgo/native/logger"
)

type Parser struct {
	source logger.Source
	lexer  js_lexer.Lexer
}

func newParser(source logger.Source) Parser {
	return Parser{
		source: source,
		lexer: js_lexer.NewLexer(source),
	};
}

func (parser * Parser) parseTopLevel() {

}
