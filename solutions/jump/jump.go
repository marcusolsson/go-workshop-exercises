package main

import (
	"errors"
	"fmt"
)

type player struct {
	stamina int
}

func (p *player) jump() error {
	if p.stamina <= 0 {
		return errors.New("out of stamina")
	}

	p.stamina--

	return nil
}

func main() {
	p := player{
		stamina: 4,
	}

	for i := 0; i < 10; i++ {
		err := p.jump()
		if err != nil {
			fmt.Println("player too tired")
		}
	}
}
