package dom_test

import (
	"testing"

	"github.com/alisdairrankine/nevis/dom"
	"github.com/alisdairrankine/nevis/dom/html"
)

func TestRenderTagToString(t *testing.T) {
	div := html.A(
		dom.Prop("href", "/"),

		html.Span("Click Here!"),
	)

	vdom := &dom.VirtualDom{}
	vdom.SetRootNode(div)
	j := dom.NewStringRenderer(false)

	str := j.RenderDom(vdom)
	if "<a href=\"/\"><span>Click Here!</span></a>" != str {
		t.Logf("unexpected output: '%s'", str)
		t.Fail()
	}
}
