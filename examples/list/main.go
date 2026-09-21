package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/list"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Track struct {
	Title  string
	Artist string
}

// example renderer for the rows
type trackDelegate struct{}

func (trackDelegate) Render(item Track, selected bool) string {
	line := fmt.Sprintf("%s — %s", item.Title, item.Artist)
	if selected {
		return lipgloss.NewStyle().Bold(true).Render(line)
	}
	return line
}

type root struct {
	list list.Model[Track]
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
		r.list.SetSize(msg.Width, msg.Height)
		return r, nil
	}

	var cmd tea.Cmd
	r.list, cmd = r.list.Update(msg)
	return r, cmd
}

func (r root) View() string {
	return r.list.View()
}

func main() {
	items := []Track{
		{Title: "Kid A", Artist: "Radiohead"},
		{Title: "Blonde", Artist: "Frank Ocean"},
		{Title: "To Pimp a Butterfly", Artist: "Kendrick Lamar"},
		{Title: "In Rainbows", Artist: "Radiohead"},
		{Title: "Channel Orange", Artist: "Frank Ocean"},
	}

	m := root{
		list: list.New(items, trackDelegate{}, list.WithTheme(theme.Default())),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
