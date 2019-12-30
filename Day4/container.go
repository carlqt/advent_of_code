package main

type Container []int

func (c *Container) store(num int) {
	if len(*c) == 0 || c.last() == num {
		*c = append(*c, num)
	} else {
		*c = nil
		*c = append(*c, num)
	}
}

func (c Container) last() int {
	if len(c) == 0 {
		return 0
	}

	return c[len(c)-1]
}

func (c Container) valid() bool {
	if len(c) == 2 {
		return true
	}

	return false
}
