package js_parser

import "github.com/hanayashiki/esgo/native/js_lexer"

func (parser *Parser) parseStatement() {
	token := parser.lexer.PeekToken()

	switch token {
	case js_lexer.TImport:
		parser.lexer.Next()
		parser.parseImport()
	}
}

// Reference: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import
func (parser *Parser) parseImport() {
	token := parser.lexer.PeekToken()

	switch token {
	// import defaultExport from "module-name";
	case js_lexer.TIdentifier:
		
	}
}
