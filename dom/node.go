package dom

import "fmt"

type Node interface{}

type ElementNode struct {
	Name          string
	EventHandlers map[string][]EventHandler
	Properties    map[string]string
	Children      []Node
}

func createElement(name string, properties map[string]string, eventHandlers map[string][]EventHandler, children ...Node) *ElementNode {
	return &ElementNode{
		Name:          name,
		Properties:    properties,
		EventHandlers: eventHandlers,
		Children:      children,
	}
}

type NodeOrProperty interface{}

type property struct {
	Key   string
	Value string
}

func Prop(Key, Value string) *property {
	return &property{
		Key:   Key,
		Value: Value,
	}
}

//Ignore components from element creation. Only care about text and basic elements
func ElementCreator(name string) func(children ...NodeOrProperty) *ElementNode {
	return func(children ...NodeOrProperty) *ElementNode {
		properties := make(map[string]string)
		eventHandlers := make(map[string][]EventHandler)
		kids := make([]Node, 0)
		for _, child := range children {
			switch child.(type) {
			case *property:
				prop := child.(*property)
				properties[prop.Key] = prop.Value
			case *ElementNode:
				kids = append(kids, child)
			case string:
				kids = append(kids, child)
			case Component:
				kids = append(kids, child.(Component).Render())
			case RegisteredEventHandler:
				if list, exists := eventHandlers[child.(RegisteredEventHandler).eventIdentifier]; exists {
					eventHandlers[child.(RegisteredEventHandler).eventIdentifier] = append(list, child.(RegisteredEventHandler).eventHandler)
				} else {
					eventHandlers[child.(RegisteredEventHandler).eventIdentifier] = []EventHandler{
						child.(RegisteredEventHandler).eventHandler,
					}
				}
			}
		}
		return createElement(name, properties, eventHandlers, kids...)
	}
}

// diffNode returns the addresses of changed nodes for incremental updates.
// Realistically, this is of no use when rendering to strings,but will allow speeding up
// redraws within a DOM environment
func diffNode(oldNode, newNode Node, address string) []string {
	diffs := make([]string, 0)
	if newNode == nil || oldNode == nil {
		diffs = append(diffs, address)
	}
	switch oldNode.(type) {
	case string:
		switch newNode.(type) {
		case string:
			if oldNode.(string) != newNode.(string) {
				diffs = append(diffs, address)
			}
		case *ElementNode:
			diffs = append(diffs, address)
		case Component:
			diffs = append(diffs, address)
		}
	case *ElementNode:
		switch newNode.(type) {
		case string:
			diffs = append(diffs, address)
		case Component:
			diffs = append(diffs, address)
		case *ElementNode:
			if newNode.(*ElementNode).Name != oldNode.(*ElementNode).Name {
				diffs = append(diffs, address)
			} else {
				if !mapIsSame(newNode.(*ElementNode).Properties, oldNode.(*ElementNode).Properties) {
					diffs = append(diffs, address)
				} else {
					//This needs to be cleverer, as removing a child will cause the whole thing to redraw
					if len(oldNode.(*ElementNode).Children) != len(newNode.(*ElementNode).Children) {
						diffs = append(diffs, address)
					} else {
						for i, oldKid := range oldNode.(*ElementNode).Children {
							addr := address + fmt.Sprintf(".%d", i)
							newKid := newNode.(*ElementNode).Children[i]
							diffs = append(diffs, diffNode(oldKid, newKid, addr)...)
						}

					}
				}
			}

		}
	case Component:
		fmt.Println(address)
		diffs = append(diffs, address)
	}
	return diffs
}
