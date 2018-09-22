package dom

import "fmt"

type StringRenderer struct {
	includeElementAddresses bool
}

func NewStringRenderer(includeElementAddresses bool) *StringRenderer {
	return &StringRenderer{
		includeElementAddresses: includeElementAddresses,
	}
}

func (r *StringRenderer) renderNodeToString(n Node, address string) string {
	switch n.(type) {
	case string:
		return n.(string)
	case *ElementNode:
		return r.renderElement(n.(*ElementNode), address)
	}
	return ""
}

func (r *StringRenderer) RenderDom(d *VirtualDom) string {
	return r.renderNodeToString(d.RootNode, "0")
}

func (r *StringRenderer) renderElement(e *ElementNode, address string) string {
	propString := ""
	for propName, prop := range e.Properties {
		propString += fmt.Sprintf(" %s=\"%s\"", propName, prop)
	}

	if r.includeElementAddresses {
		propString += fmt.Sprintf(" data-address-id=\"%s\"", address)
	}
	if len(e.Children) == 0 {
		return fmt.Sprintf("<%s%s/>", e.Name, propString)
	}
	contents := ""
	for i, child := range e.Children {
		contents += r.renderNodeToString(child, fmt.Sprintf("%s.%d", address, i))
	}
	return fmt.Sprintf("<%s%s>%s</%s>", e.Name, propString, contents, e.Name)
}
