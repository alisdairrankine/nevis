package main

import (
	"fmt"

	"github.com/alisdairrankine/nevis/dom"
	"github.com/alisdairrankine/nevis/dom/html"
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
	j := dom.NewStringRenderer(true)

	fmt.Println(j.RenderDom(vdom))
}
