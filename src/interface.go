package main

import "fmt"

type Moviea struct {
	Name string
	Rating float32
}
// 函数 丢失了绑定特性
//func summary(m *Moviea) string {
//	//
//	fmt.Println(m.Name)
//	return m.Name
//}

func (m *Moviea)summary() string{
	fmt.Println(m.Name)
	return m.Name
}
// 接口func后面添加另一个参数部分，用于接受单个参数

func main()  {
	m := Moviea{
		Name:   "zhang",
		Rating: 0,
	}
	//summary(&m)
	m.summary()
}