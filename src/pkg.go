package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	)

func main(){
	data := [][]string{
		[]string {"zhang1","10","10/20","(1.3, 2.3,3.4)"},
		[]string {"zhang2","10","10/20","(1.3, 2.3,3.4)"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPC","Speed","Power","Location"})
	table.AppendBulk(data)
	table.Render()
}