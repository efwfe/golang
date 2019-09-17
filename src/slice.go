package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "aaa"
	cheeses[1] = "bbbb"
	cheeses = append(cheeses, "cccc")
	fmt.Println(cheeses)
}
