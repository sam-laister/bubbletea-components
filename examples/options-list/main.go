package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/checkbox"
	optionslist "github.com/sam-laister/sam-laister-bubbletea-components-library/options-list"
	textinput "github.com/sam-laister/sam-laister-bubbletea-components-library/text-input"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type root struct {
	options optionslist.Model
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
	items := []optionslist.Field{
		checkbox.New("Kid A", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("Blonde", false, checkbox.WithTheme(theme.Default())),
		checkbox.New("To Pimp a Butterfly", false, checkbox.WithTheme(theme.Default())),
		textinput.New("In Rainbows", "10", textinput.WithTheme(theme.Default())),
		textinput.New("Channel Orange", "5", textinput.WithTheme(theme.Default())),
	}

	m := root{
		options: optionslist.New(items, theme.Default()),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
