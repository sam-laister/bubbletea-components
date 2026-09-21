package list

type ItemDelegate[T any] interface {
	Render(item T, selected bool) string
}
