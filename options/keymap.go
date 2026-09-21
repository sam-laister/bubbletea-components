package options

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Activate key.Binding
	Cancel   key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Activate: key.NewBinding(key.WithKeys("enter", " ")),
		Cancel:   key.NewBinding(key.WithKeys("esc")),
	}
}
