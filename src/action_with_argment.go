package main

import "fmt"

type Triangle struct {
	height float64
	width float64
}
// * 不加 会在接口中创建副本 导致无法修改原来的结构体
func (t *Triangle)changeBase(data float64)  {
	t.height = data
}

func main()  {
	t := Triangle{
		height: 10,
		width:  20,
	}
	t.changeBase(12)
	fmt.Println(t.height)
}