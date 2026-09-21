package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	optionslist "github.com/sam-laister/sam-laister-bubbletea-components-library/options-list"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Model struct {
	input  textinput.Model
	label  string
	styles Styles
}

func New(label, value string, opts ...Option) *Model {
	cfg := config{
		styles: DefaultStyles(theme.Default()),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	input := textinput.New()
	input.SetValue(value)

	return &Model{
		input:  input,
		label:  label,
		styles: cfg.styles,
	}
}

func (m *Model) Blur() {
	m.input.Blur()
}

func (m *Model) CapturesInput() bool {
	return true
}

func (m *Model) Focus() {
	m.input.Focus()
}

func (m *Model) Focused() bool {
	return m.input.Focused()
}

func (m *Model) Update(msg tea.Msg) (optionslist.Field, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	labelStyle := m.styles.Label
	if m.input.Focused() {
		labelStyle = m.styles.LabelFocused
	}

	return labelStyle.Render(m.label) + " " + m.input.View()
}
