package main

import "fmt"

// STRUCTS AND INTERFACE GOES HERE

func battle(c warcrier) {
	fmt.Println(p.warcry())
}

func main() {
	p := player{
		nickname: "Anonymous",
	}
	e := enemy{
		name: "Yog-Sothoth",
	}

	battle(p)
	battle(e)
}
