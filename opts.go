package counter

type Option func(c *Counter)

func SetupCounter(names ...string) Option {
	return func(c *Counter) {
		for _, n := range names {
			if _, has := c.counting[n]; !has {
				c.counting[n] = 0
			}
		}
	}
}
