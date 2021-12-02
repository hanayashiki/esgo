package js_lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hanayashiki/esgo/native/test_utils"
)

func TestIdentifier(t * testing.T) {
	source := test_utils.SourceFromText("foo")
	lexer := NewLexer(source)

	assert.Equal(t, lexer.currentIdentifierName, "foo", "Should parse simple identifier")

	source = test_utils.SourceFromText("  foo")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentIdentifierName, "foo", "Should parse simple identifier after simple whitespace")

	source = test_utils.SourceFromText("\n \tfoo")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentIdentifierName, "foo", "Should parse simple identifier after whitespace series")

	source = test_utils.SourceFromText("香港")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentIdentifierName, "香港", "Should parse unicode identifier")

	source = test_utils.SourceFromText("  𠄯𠄯𠄯𠄯")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentIdentifierName, "𠄯𠄯𠄯𠄯", "Should parse unicode identifier")

	source = test_utils.SourceFromText("  foo foooo")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentIdentifierName, "foooo", "Should parse continuous identifiers")

	source = test_utils.SourceFromText(":香港")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentIdentifierName, "香港", "Should parse unicode identifier after delimiter")
}

func TestDelimiters(t * testing.T) {
	source := test_utils.SourceFromText("{")
	lexer := NewLexer(source)

	assert.Equal(t, lexer.currentToken, TOpenBrace, "Should parse simple delimiters")

	source = test_utils.SourceFromText("}")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentToken, TCloseBrace, "Should parse simple delimiters")

	source = test_utils.SourceFromText("     }")
	lexer = NewLexer(source)

	assert.Equal(t, lexer.currentToken, TCloseBrace, "Should parse simple delimiters after whitespace")

	source = test_utils.SourceFromText("fooooooo    }")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TCloseBrace, "Should parse simple delimiters after identifier")

	source = test_utils.SourceFromText("@    }foo{")
	lexer = NewLexer(source)
	lexer.Next()
	lexer.Next()
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TOpenBrace, "Should parse multiple simple delimiters")

	source = test_utils.SourceFromText("a ? b : c")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TQuestion, "Should parse TQuestion")

	source = test_utils.SourceFromText("a ?? b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TQuestionQuestion, "Should parse TQuestionQuestion")

	source = test_utils.SourceFromText("a ??= b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TQuestionQuestionEquals, "Should parse TQuestionQuestionEquals")

	source = test_utils.SourceFromText("a?.[b]")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TQuestionDot, "Should parse TQuestionDot")

	source = test_utils.SourceFromText("a % b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TPercent, "Should parse TPercent")

	source = test_utils.SourceFromText("a %= b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TPercentEquals, "Should parse TPercentEquals")

	source = test_utils.SourceFromText("a ^ b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TCaret, "Should parse TCaret")

	source = test_utils.SourceFromText("a ^= b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TCaretEquals, "Should parse TCaretEquals")

	source = test_utils.SourceFromText("a & b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TAmpersand, "Should parse TAmpersand")

	source = test_utils.SourceFromText("a && b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TAmpersandAmpersand, "Should parse TAmpersandAmpersand")

	source = test_utils.SourceFromText("a &&= b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TAmpersandAmpersandEquals, "Should parse TAmpersandAmpersandEquals")

	source = test_utils.SourceFromText("a | b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TBar, "Should parse TBar")

	source = test_utils.SourceFromText("a || b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TBarBar, "Should parse TBarBar")

	source = test_utils.SourceFromText("a ||= b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TBarBarEquals, "Should parse TBarBarEquals")

	source = test_utils.SourceFromText("a + b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TPlus, "Should parse TPlus")

	source = test_utils.SourceFromText("a += b")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TPlusEquals, "Should parse TPlusEquals")

	source = test_utils.SourceFromText("a++")
	lexer = NewLexer(source)
	lexer.Next()

	assert.Equal(t, lexer.currentToken, TPlusPlus, "Should parse TPlusPlus")
}