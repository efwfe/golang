package basic

import "fmt"

const bollingF = 212.0

func Biling(){
	var f = bollingF
	var c = (f -32) * 5/9
	fmt.Printf("boling point = %gF or %gC \n", f,c)
	// 输出 boling point = 212F or 100C
	Okay()
}
