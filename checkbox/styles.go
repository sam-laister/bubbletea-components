package checkbox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Styles struct {
	Checked      lipgloss.Style
	Unchecked    lipgloss.Style
	Label        lipgloss.Style
	LabelFocused lipgloss.Style
}

func DefaultStyles(t theme.Theme) Styles {
	return Styles{
		Checked: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),
		Unchecked: lipgloss.NewStyle().
			Foreground(t.Muted),
		Label: lipgloss.NewStyle().
			Foreground(t.Text),
		LabelFocused: lipgloss.NewStyle().
			Foreground(t.BorderFocus),
	}
}
