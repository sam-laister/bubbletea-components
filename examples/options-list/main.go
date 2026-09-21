package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/checkbox"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/options"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type root struct {
	options options.Model
}

func (r root) Init() tea.Cmd {
	return nil
}

func (r root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return r, tea.Quit
		}
	case tea.WindowSizeMsg:
		r.options.SetSize(msg.Width, msg.Height)
		return r, nil
	}

	var cmd tea.Cmd
	r.options, cmd = r.options.Update(msg)
	return r, cmd
}

func (r root) View() string {
	return r.options.View()
}

func main() {
	items := []options.Field{
		checkbox.New("Kid A", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("Blonde", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("To Pimp a Butterfly", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("In Rainbows", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("Channel Orange", false, checkbox.WithTheme(theme.Default())),
	}

	m := root{
		options: options.New(items, theme.Default()),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
