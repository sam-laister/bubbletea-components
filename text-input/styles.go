package textinput

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Styles struct {
	Label        lipgloss.Style
	LabelFocused lipgloss.Style
}

func DefaultStyles(t theme.Theme) Styles {
	return Styles{
		Label: lipgloss.NewStyle().
			Foreground(t.Text),
		LabelFocused: lipgloss.NewStyle().
			Foreground(t.BorderFocus),
	}
}
