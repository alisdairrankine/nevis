package main

import (
	"fmt"

	"github.com/alisdairrankine/nevis/dom"
	"github.com/alisdairrankine/nevis/dom/html"
	"github.com/alisdairrankine/nevis/renderer/string"
)

func main() {

	div := html.Div(
		dom.Prop("class", "header"),

		html.H1("Hello"),
		html.A(
			dom.Prop("href", "/"),

			html.Span("Click here!"),
		),
	)

	vdom := &dom.VirtualDom{}
	vdom.SetRootNode(div)
	j := string.NewStringRenderer(true)

	fmt.Println(j.RenderDom(vdom))
}
