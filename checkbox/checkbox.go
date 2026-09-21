package checkbox

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	optionslist "github.com/sam-laister/sam-laister-bubbletea-components-library/options-list"
	"github.com/sam-laister/sam-laister-bubbletea-components-library/theme"
)

type Model struct {
	label   string
	checked bool
	focused bool
	keyMap  KeyMap
	styles  Styles
}

func New(label string, checked bool, opts ...Option) *Model {
	cfg := config{
		styles: DefaultStyles(theme.Default()),
		keyMap: DefaultKeyMap(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Model{
		label:   label,
		checked: checked,
		focused: false,
		keyMap:  cfg.keyMap,
		styles:  cfg.styles,
	}
}

func (c *Model) Blur() {
	c.focused = false
}

func (c *Model) CapturesInput() bool {
	return false
}

func (c *Model) Focus() {
	c.focused = true
}

func (c *Model) Focused() bool {
	return c.focused
}

func (c *Model) Update(msg tea.Msg) (optionslist.Field, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keyMap.Toggle):
			c.checked = !c.checked
		}
	}

	return c, cmd
}

func (c *Model) View() string {
	glyph := "[ ]"
	glyphStyle := c.styles.Unchecked
	if c.checked {
		glyph = "[x]"
		glyphStyle = c.styles.Checked
	}

	labelStyle := c.styles.Label
	if c.focused {
		labelStyle = c.styles.LabelFocused
	}

	return glyphStyle.Render(glyph) + " " + labelStyle.Render(c.label)
}

func (c *Model) Checked() bool {
	return c.checked
}

var _ optionslist.Field = (*Model)(nil)
