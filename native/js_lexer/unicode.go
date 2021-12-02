package js_lexer

func isIdentifierStart(code rune) bool {
	if code == '_' || code == '$' || (code >= 'a' && code <= 'z') || (code >= 'A' && code <= 'Z') {
		return true
	}

	if code <= 0x7f {
		return false
	}

	return code != -1 // Don't bother to handle strange unicode cases
}

func isIdentifierPart(code rune) bool {
	return code != -1 && (isIdentifierStart(code) || (code >= '0' && code <= '9'))
}

func isWhitespace(code rune) bool {
	switch code {
	case
		'\u0009', // character tabulation
		'\u000A', // \n
		'\u000B', // line tabulation
		'\u000C', // form feed
		'\u000D', // \r
		'\u0020', // space
		'\u00A0', // no-break space

		// Unicode "Space_Separator" code points
		'\u1680', // ogham space mark
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000', // ideographic space

		'\uFEFF': // zero width non-breaking space
		return true

	default:
		return false
	}
}