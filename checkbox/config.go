package checkbox

import "github.com/sam-laister/sam-laister-bubbletea-components-library/theme"

type config struct {
	styles Styles
	keyMap KeyMap
}

type Option func(*config)

func WithTheme(t theme.Theme) Option {
	return func(c *config) {
		c.styles = DefaultStyles(t)
	}
}

func WithKeyMap(km KeyMap) Option {
	return func(c *config) {
		c.keyMap = km
	}
}
