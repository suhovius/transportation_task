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
		// TODO: rename resultTable to deliveryTable
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
	fmt.Printf("Delivery Cost: %d", task.deliveryCost())
	fmt.Println()

	// Transport potentials method starts here

	// 1 Size distribution sum check

	// 2 Degeneracy check
	// http://cyclowiki.org/wiki/%D0%92%D1%8B%D1%80%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%BD%D0%BE%D1%81%D1%82%D1%8C_%D0%B2_%D1%82%D1%80%D0%B0%D0%BD%D1%81%D0%BF%D0%BE%D1%80%D1%82%D0%BD%D0%BE%D0%B9_%D0%B7%D0%B0%D0%B4%D0%B0%D1%87%D0%B5
	// And might need to have added small floating point numbers here (0.001) to fix this degeneracy issue
}
