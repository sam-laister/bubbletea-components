package checkbox

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Toggle key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Toggle: key.NewBinding(
			key.WithKeys("enter", "space"),
			key.WithHelp("enter", "select"),
		),
	}
}
