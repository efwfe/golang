package main

import "fmt"

func main() {

	var cheeses = make([]string, 3)
	cheeses[0] = "aaa"
	cheeses[1] = "bbb"
	cheeses[2] = "ccc"
	fmt.Println(cheeses)
	cheeses = append(cheeses[:3], cheeses[1:]...)
	var ano = make([]string, 8)
	copy(ano, cheeses)
	fmt.Println(cheeses)
	fmt.Println(ano)
	fmt.Println(len(ano))
}
