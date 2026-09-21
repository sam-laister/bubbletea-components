package optionslist

type config struct {
	keyMap KeyMap
}

type Option func(*config)

func WithKeyMap(km KeyMap) Option {
	return func(c *config) {
		c.keyMap = km
	}
}
