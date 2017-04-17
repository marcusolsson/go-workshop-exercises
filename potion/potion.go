package main

import "fmt"

// STRUCT GOES HERE

func main() {
	p := player{
		nickname: "Anonymous",
		health:   12,
	}

	p.consumePotion()

	if p.health == 100 {
		fmt.Println(p.nickname, "lives to fight another day")
	} else {
		fmt.Println(p.nickname, "still suffers")
	}
}
