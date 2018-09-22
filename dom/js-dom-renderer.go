package dom

import (
	"fmt"
	"strconv"
	"strings"

	jsdom "honnef.co/go/js/dom"
)

type JSDomRenderer struct {
	vDom *VirtualDom
}

func (r *JSDomRenderer) RenderDom() {
	parent := jsdom.GetWindow().Document().GetElementsByTagName("body")
	if len(parent) < 1 {
		panic("I gots no body")
	}
	r.renderVirtualElement(parent[0], r.vDom.RootNode, "0", false)
}

func (r *JSDomRenderer) renderVirtualElement(parent jsdom.Node, node Node, address string, replace bool) {
	var n Node
	switch node.(type) {
	case Component:
		n = node.(Component).Render()
	default:
		n = node
	}
	if n == nil {
		return
	}
	switch n.(type) {
	case string:
		fmt.Println("string!!")
		domel := jsdom.GetWindow().Document().CreateTextNode(n.(string))

		if replace {
			parent.ParentElement().ReplaceChild(domel, parent)
		} else {
			parent.AppendChild(domel)
		}
	case Component:
		componentChildren := n.(Component).Render()
		fmt.Println(componentChildren, address)
		r.renderVirtualElement(parent, componentChildren, address, replace)

	case *ElementNode:
		fmt.Println("node!!")
		el := n.(*ElementNode)
		domel := jsdom.GetWindow().Document().CreateElement(el.Name)

		for prop, value := range el.Properties {
			domel.SetAttribute(prop, value)
		}
		for event, handlers := range el.EventHandlers {
			domel.AddEventListener(event, true, func(ev jsdom.Event) {
				//create synthetic event
				e := Event{}
				for _, handler := range handlers {
					handler(e)
				}
			})
		}
		domel.SetAttribute("loom-ui-id", address)
		if replace {
			parent.ParentElement().ReplaceChild(domel, parent)
		} else {
			parent.AppendChild(domel)
		}
		for i, child := range el.Children {
			r.renderVirtualElement(domel, child, fmt.Sprintf("%s.%d", address, i), replace)
		}
	default:
		fmt.Printf("unknown node: %+v", n)
	}
}

func getChildFromRelativeAddress(parent jsdom.Node, address string) jsdom.Node {
	splitAddress := strings.Split(address, ".")

	index, _ := strconv.Atoi(splitAddress[0])
	children := parent.ChildNodes()
	if len(splitAddress) == 1 {
		return children[index]
	}
	newAddress := strings.Join(splitAddress[1:], ".")
	return getChildFromRelativeAddress(children[index], newAddress)

}

func (r *JSDomRenderer) getDomElementByAddress(address string) jsdom.Node {
	body := jsdom.GetWindow().Document().GetElementsByTagName("body")[0]
	return getChildFromRelativeAddress(body, address)
}

func (r *JSDomRenderer) Rerender(addresses []string) {

	for _, address := range addresses {
		node := r.vDom.GetNodeByAddress(address)
		el := r.getDomElementByAddress(address)

		if el != nil {
			r.renderVirtualElement(el, node, address, true)
		}
	}

}

func NewJSDomRenderer(vDom *VirtualDom) *JSDomRenderer {
	r := &JSDomRenderer{vDom: vDom}
	vDom.listenToUpdates(func(addresses []string) {
		r.Rerender(addresses)
	})
	return r
}
