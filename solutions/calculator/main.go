package main

import (
	"fmt"

	"github.com/marcusolsson/go-workshop-exercises/solutions/calculator/calc"
)

func main() {
	x := 7
	y := 5

	sum := calc.Add(x, y)

	fmt.Println(sum)
}
