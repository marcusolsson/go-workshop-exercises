package greeter

type greeter struct {
	name string
}

func (g greeter) greet() string {
	return "Hello, " + g.name + "!"
}
