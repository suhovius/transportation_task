package main

import "fmt"

// TaskState contains transportation task state parameters
type taskState struct {
	supplyList  []int
	demandList  []int
	costTable   [][]int
	resultTable [][]int
}

func pritnArrayOfArrays(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

// Print prints inner state of the taskState structure
func (s taskState) print() {
	fmt.Printf("supplyList = %v\n", s.supplyList)
	fmt.Printf("demandList = %v\n", s.demandList)
	fmt.Println("costTable =")
	pritnArrayOfArrays(s.costTable)
	fmt.Println("resultTable =")
	pritnArrayOfArrays(s.resultTable)
}
