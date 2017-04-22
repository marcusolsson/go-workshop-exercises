package main

type player struct {
	stamina int
}

func (p *player) jump() {
	// ERROR IF NO STAMINA LEFT
	p.stamina--
}

func main() {
	p := player{
		stamina: 4,
	}

	for i := 0; i < 10; i++ {
		p.jump()
		// HANDLE ERRORS
	}
}
