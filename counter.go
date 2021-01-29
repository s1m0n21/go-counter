package counter

import "sync"

type Counter struct {
	lk sync.RWMutex

	counting map[string]int
}

func New(opts ...Option) *Counter {
	c := &Counter{
		lk:       sync.RWMutex{},
		counting: make(map[string]int),
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *Counter) Plus(name string, count int) {
	c.lk.Lock()
	defer c.lk.Unlock()

	if _, has := c.counting[name]; has {
		c.counting[name] += count
		return
	}

	c.counting[name] = count
}

func (c *Counter) Minus(name string, count int) {
	c.lk.Lock()
	defer c.lk.Unlock()

	if _, has := c.counting[name]; has {
		c.counting[name] -= count
		return
	}
}

func (c *Counter) Count(name string) (int, bool) {
	c.lk.RLock()
	defer c.lk.RUnlock()

	count, has := c.counting[name]
	return count, has
}
