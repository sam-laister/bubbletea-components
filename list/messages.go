package list

import tea "github.com/charmbracelet/bubbletea"

type CursorMsg struct {
	Index int
}

func CursorCmd(index int) tea.Cmd {
	return func() tea.Msg {
		return CursorMsg{
			Index: index,
		}
	}
}

type SelectMsg[T any] struct {
	Index int
	Item  T
}

func SelectCmd[T any](index int, item T) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg[T]{
			Index: index,
			Item:  item,
		}
	}
}
