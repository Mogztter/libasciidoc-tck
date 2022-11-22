package main

import (
	"bufio"
	"fmt"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/mogztter/libasciidoc-tck/pkg"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	//output := bufio.NewWriter(os.Stdout)
	doc, _ := pkg.Parse(input, configuration.NewConfiguration(configuration.WithHeaderFooter(true)))
	for _, element := range doc.Elements {
		v, ok := element.(*types.Section)
		if ok {
			fmt.Printf("%i", v.Level)
			fmt.Printf("%s", v.Title)
		}

		fmt.Printf("%v\n", v)
		fmt.Printf("%v\n", element)
		fmt.Printf("%T\n", element)
	}
}
