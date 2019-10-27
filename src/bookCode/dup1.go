// 找出重复行 输出标准输入出现次数大于1的行 前面的是次数
package main

import (
	"bufio" // 处理输入与输出
	"fmt"
	"os"
)

func main(){
	counts := make(map[string] int)
	input := bufio.NewScanner(os.Stdin) // 读取输入 以行或者单词为单位断开
	for input.Scan(){
		if input.Text() == "quit" {
			break;
		}

		counts[input.Text()] ++
	}

	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%n", n, line)
		}

	}
}