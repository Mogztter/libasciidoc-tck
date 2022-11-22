package pkg

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"io"
)

func Parse(source io.Reader, config *configuration.Configuration) (*types.Document, error) {
	return parser.ParseDocument(source, config)
}
