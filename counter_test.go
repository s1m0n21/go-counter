package counter

import "testing"

func TestCounter(t *testing.T) {
	c := New(SetupCounter("A", "B", "C"))

	c.Plus("A", 1)
	a, has := c.Count("A")
	t.Logf("A: count: %d, has: %v", a, has)

	c.Minus("A", 1)
	a, has = c.Count("A")
	t.Logf("A: count: %d, has: %v", a, has)

	c.Minus("D", 1)
	d, has := c.Count("D")
	t.Logf("D: count: %d, has: %v", d, has)

	c.Plus("F", 1)
	f, has := c.Count("F")
	t.Logf("F: count: %d, has: %v", f, has)
}
