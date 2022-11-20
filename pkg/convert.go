package pkg

import (
	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"io"
)

func ConvertToASG(source io.Reader, output io.Writer, config *configuration.Configuration) (types.Metadata, error) {
	return libasciidoc.Convert(source, output, config)
}
