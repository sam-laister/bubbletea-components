package list

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Styles struct {
	Container lipgloss.Style
	Cursor    lipgloss.Style
	Empty     lipgloss.Style
}

func DefaultStyles(t theme.Theme) Styles {
	return Styles{
		Container: lipgloss.NewStyle(),
		Cursor: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),
		Empty: lipgloss.NewStyle().
			Foreground(t.Muted).
			Italic(true),
	}
}
