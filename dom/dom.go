package dom

import (
	"strconv"
	"strings"
)

// WARNING: EXPERIMENTAL Virtual DOM

type VirtualDom struct {
	RootNode  Node
	listeners []func([]string)
}

func (d *VirtualDom) SetRootNode(node Node) {
	switch node.(type) {
	case string:
		d.RootNode = node
	case *ElementNode:
		d.RootNode = node
	case Component:
		d.RootNode = node
	default:
		panic("invalid dom node")
	}
}

func (d *VirtualDom) listenToUpdates(listener func([]string)) {
	if d.listeners == nil {
		d.listeners = make([]func([]string), 0)
	}
	d.listeners = append(d.listeners, listener)
}

func (d *VirtualDom) GetNodeByAddress(address string) Node {
	if address == "0" {
		return d.RootNode
	}
	addressParts := strings.Split(address, ".")
	return GetDescendentByRelativeAddress(d.RootNode, strings.Join(addressParts[1:], "."))
}

func GetDescendentByRelativeAddress(n Node, address string) Node {
	switch n.(type) {
	case string:
		return nil
	case *ElementNode:
		for i, child := range n.(*ElementNode).Children {
			addressParts := strings.Split(address, ".")
			j, _ := strconv.Atoi(addressParts[0])
			if j == i {
				if len(address) > 1 {
					addr := strings.Join(addressParts[1:], ".")
					return GetDescendentByRelativeAddress(child, addr)
				}
				return child

			}
		}
	case Component:
		child := n.(Component).Render()
		return GetDescendentByRelativeAddress(child, address)

	}
	return nil
}

func (d *VirtualDom) UpdateRootNode(newNode Node) {
	//calculate diffs in tree, replace new tree.
	updates := diffNode(d.RootNode, newNode, "0")
	d.RootNode = newNode
	for _, listener := range d.listeners {
		listener(updates)
	}

}

func mapIsSame(old, new map[string]string) bool {
	same := true
	if len(old) != len(new) {
		same = false
	} else {
		for key, value := range old {
			if newValue, exists := new[key]; exists && value != newValue {
				same = false
			}
		}
	}
	return same
}
