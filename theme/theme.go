package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Accent      lipgloss.AdaptiveColor
	Border      lipgloss.AdaptiveColor
	BorderFocus lipgloss.AdaptiveColor
	Text        lipgloss.AdaptiveColor
	Muted       lipgloss.AdaptiveColor
}

func Default() Theme {
	return Theme{
		Accent:      lipgloss.AdaptiveColor{Light: "#0F6F8C", Dark: "#38C9F0"},
		Border:      lipgloss.AdaptiveColor{Light: "#C7D0DA", Dark: "#2C3542"},
		BorderFocus: lipgloss.AdaptiveColor{Light: "#0F6F8C", Dark: "#38C9F0"},
		Text:        lipgloss.AdaptiveColor{Light: "#1B2430", Dark: "#DCE3EA"},
		Muted:       lipgloss.AdaptiveColor{Light: "#5B6A7A", Dark: "#8B96A5"},
	}
}
