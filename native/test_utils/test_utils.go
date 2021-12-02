package test_utils

import (
	"github.com/hanayashiki/esgo/native/logger"
)

func SourceFromText(contents string) logger.Source {
	return logger.Source{
		Index: 0,
		KeyPath:        logger.Path{Text: "<stdin>"},
		PrettyPath:     "<stdin>",
		Contents:       contents,
		IdentifierName: "stdin",
	}
}