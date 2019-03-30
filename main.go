package main

import (
	"fmt"
)

type taskState struct {
	supplyList  []int
	demandList  []int
	costTable   [][]int
	resultTable [][]int
}

func main() {
	state := taskState{
		supplyList: []int{30, 40, 20},
		demandList: []int{20, 30, 30, 10},
		costTable: [][]int{
			{2, 3, 2, 4},
			{3, 2, 5, 1},
			{4, 3, 2, 6},
		},
	}

	state.print()
}

func pritnArrayOfArrays(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func (s taskState) print() {
	fmt.Printf("supplyList = %v\n", s.supplyList)
	fmt.Printf("demandList = %v\n", s.demandList)
	fmt.Println("costTable =")
	pritnArrayOfArrays(s.costTable)
	pritnArrayOfArrays(s.resultTable)
}
