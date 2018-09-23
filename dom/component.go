package dom

// Component It's all been building up to this: the component!
type Component interface {
	HasUpdated() bool
	Render() Node
}

// PureComponent is embeddable and provides a non-updatable component
type PureComponent struct{}

// HasUpdated always returns false, as PureComponents never change
func (c *PureComponent) HasUpdated() bool { return false }
