package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/mogztter/libasciidoc-tck/pkg"
	"os"
)

type AbstractSemanticGraph types.Document
type Paragraph types.Paragraph

func (p Paragraph) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	elements := p.Elements
	if s, ok := elements[0].(*types.StringElement); ok {
		m["type"] = "Paragraph"
		m["lines"] = []string{s.Content}
	}
	return json.Marshal(m)
}

func (p AbstractSemanticGraph) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	var elements []interface{}
	for _, element := range p.Elements {
		if _, isAttributeDeclaration := element.(*types.AttributeDeclaration); isAttributeDeclaration {
			continue
		}
		if para, isParagraph := element.(*types.Paragraph); isParagraph {
			elements = append(elements, Paragraph{
				Attributes: para.Attributes,
				Elements:   para.Elements,
			})
		} else {
			elements = append(elements, element)
		}
	}
	m["body"] = elements
	return json.Marshal(m)
}

func main() {
	input := bufio.NewReader(os.Stdin)
	//output := bufio.NewWriter(os.Stdout)
	doc, _ := pkg.Parse(input, configuration.NewConfiguration(configuration.WithHeaderFooter(true)))
	asg := AbstractSemanticGraph{
		Elements:          doc.Elements,
		ElementReferences: doc.ElementReferences,
		Footnotes:         doc.Footnotes,
		TableOfContents:   doc.TableOfContents,
	}
	byteSlice, _ := json.Marshal(asg)
	fmt.Println(string(byteSlice))
}
