package main

type Module struct {
	Mass int
}

func (m *Module) requiredFuel() int {
	return m.Mass/3 - 2
}

type Modules []Module

func (modules Modules) totalFuel() int {
	var total int

	for _, module := range modules {
		total = total + module.requiredFuel()
	}

	return total
}
