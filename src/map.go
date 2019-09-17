package main

import (
	"fmt"
)

func main() {
	var players = make(map[string]int)
	players["cook"] = 32
	players["worker"] = 12
	delete(players, "worker")
	fmt.Println(players)

}
