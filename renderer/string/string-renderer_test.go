package string_test

import (
	"testing"

	"github.com/alisdairrankine/nevis/dom"
	"github.com/alisdairrankine/nevis/dom/html"
	"github.com/alisdairrankine/nevis/renderer/string"
)

func TestRenderTagToString(t *testing.T) {
	div := html.A(
		dom.Prop("href", "/"),

		html.Span("Click Here!"),
	)

	vdom := &dom.VirtualDom{}
	vdom.SetRootNode(div)
	j := string.NewStringRenderer(false)

	str := j.RenderDom(vdom)
	if "<a href=\"/\"><span>Click Here!</span></a>" != str {
		t.Logf("unexpected output: '%s'", str)
		t.Fail()
	}
}
