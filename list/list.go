package list

import (
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/internal/clamp"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Model[T any] struct {
	items    []T
	cursor   int
	delegate ItemDelegate[T]
	keyMap   KeyMap
	styles   Styles
	viewport viewport.Model
}

// list.New(items, delegate, list.WithTheme(theme.Default()))
func New[T any](items []T, delegate ItemDelegate[T], opts ...Option) Model[T] {
	cfg := config{
		styles: DefaultStyles(theme.Default()),
		keyMap: DefaultKeyMap(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	return Model[T]{
		items:    items,
		delegate: delegate,
		keyMap:   cfg.keyMap,
		styles:   cfg.styles,
	}
}

func (m *Model[T]) Update(msg tea.Msg) (Model[T], tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Up):
			m.cursor = clamp.Clamp(m.cursor-1, 0, len(m.items)-1)
			cmd = CursorCmd(m.cursor)
		case key.Matches(msg, m.keyMap.Down):
			m.cursor = clamp.Clamp(m.cursor+1, 0, len(m.items)-1)
			cmd = CursorCmd(m.cursor)
		case key.Matches(msg, m.keyMap.Select):
			cmd = SelectCmd(m.cursor, m.items[m.cursor])
		}
	}

	m.syncViewport()

	return *m, cmd
}

func (m *Model[T]) View() string {
	if len(m.items) == 0 {
		return m.styles.Empty.Render("No items")
	}

	rows := make([]string, len(m.items))
	for i, item := range m.items {
		selected := i == m.cursor
		row := m.delegate.Render(item, selected)

		gutter := "  "
		if selected {
			gutter = m.styles.Cursor.Render("> ")
		}
		rows[i] = gutter + row
	}

	vp := m.viewport
	vp.SetContent(lipgloss.JoinVertical(lipgloss.Left, rows...))

	return m.styles.Container.Width(vp.Width).Render(vp.View())
}

func (m *Model[T]) SetSize(width, height int) {
	m.viewport.Width = width
	m.viewport.Height = height
}

func (m *Model[T]) SelectedItem() T {
	return m.items[m.cursor]
}

func (m *Model[T]) SetSelectedItem(item T) {
	m.items[m.cursor] = item
}

func (m *Model[T]) syncViewport() {
	height := m.viewport.Height
	switch {
	case m.cursor < m.viewport.YOffset:
		m.viewport.SetYOffset(m.cursor)
	case height > 0 && m.cursor >= m.viewport.YOffset+height:
		m.viewport.SetYOffset(m.cursor - height + 1)
	}
}
