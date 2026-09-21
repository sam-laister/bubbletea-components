package optionslist

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/list"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Field interface {
	Update(tea.Msg) (Field, tea.Cmd)
	View() string
	Focus()
	Blur()
	Focused() bool
	CapturesInput() bool
}

type Mode int

const (
	ModeNavigate Mode = iota
	ModeEdit
)

type Model struct {
	mode   Mode
	rows   list.Model[Field]
	keyMap KeyMap
}

func New(fields []Field, t theme.Theme, opts ...Option) Model {
	cfg := config{
		keyMap: DefaultKeyMap(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	return Model{
		mode:   ModeNavigate,
		rows:   list.New(fields, fieldDelegate{}, list.WithTheme(t)),
		keyMap: cfg.keyMap,
	}
}

type fieldDelegate struct{}

func (fieldDelegate) Render(f Field, selected bool) string {
	return f.View()
}

func (m *Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.mode == ModeEdit {
		return m.updateEdit(msg)
	}
	return m.updateNavigate(msg)
}

func (m *Model) updateNavigate(msg tea.Msg) (Model, tea.Cmd) {
	if k, ok := msg.(tea.KeyMsg); ok && key.Matches(k, m.keyMap.Activate) {
		field := m.rows.SelectedItem()

		if field.CapturesInput() {
			field.Focus()
			m.rows.SetSelectedItem(field)
			m.mode = ModeEdit
			return *m, nil
		}

		updated, cmd := field.Update(msg)
		m.rows.SetSelectedItem(updated)
		return *m, cmd
	}

	var cmd tea.Cmd
	m.rows, cmd = m.rows.Update(msg)
	return *m, cmd
}

func (m *Model) updateEdit(msg tea.Msg) (Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
		field := m.rows.SelectedItem()
		field.Blur()
		m.rows.SetSelectedItem(field)
		m.mode = ModeNavigate
		return *m, nil
	}
	field, cmd := m.rows.SelectedItem().Update(msg)
	m.rows.SetSelectedItem(field)
	return *m, cmd
}

func (m *Model) View() string {
	return m.rows.View()
}

func (m *Model) SetSize(width, height int) {
	m.rows.SetSize(width, height)
}
