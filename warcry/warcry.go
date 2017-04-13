package main

import "fmt"

// STRUCT GOES HERE

func main() {
	p := player{
		nickname: "Anonymous",
	}

	fmt.Println("Fearlessly,", p.nickname, "dived into the abyss, screaming:", p.warcry())
}
