package main

import "fmt"

func main() {
	task := Task{
		supplyList: []int{30, 40, 20},
		demandList: []int{20, 30, 30, 10},
		costTable: [][]int{
			{2, 3, 2, 4},
			{3, 2, 5, 1},
			{4, 3, 2, 6},
		},
		resultTable: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}

	fmt.Println("Initial State")
	task.print()
	fmt.Println()

	kind := task.performBalancing()
	switch kind {
	case "nothing":
		fmt.Println("Balancing: Task is already balanced. Skip balancing")
	case "fake_demand":
		fmt.Println("Balancing: Add fake demand")
	case "fake_supply":
		fmt.Println("Balancing: Add fake supply")
	}
	task.print()
	fmt.Println()

	fmt.Println("Base Plan: Calculated with 'North West Corner' method")
	task.northWestCorner()
	task.print()
	fmt.Println()
}
