package main

import "fmt"

type Point struct {
	X,Y int
}

type Circle struct {
	Point
	Spokes int
}

func main(){
	var c Circle
	c.X = 12
	fmt.Printf("%#",c)

}
