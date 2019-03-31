package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Remove this when frontend web server will be created
// that will handle these requests with such parameters
func testJSONData() []byte {
	inputTest := Params{
		SupplyList: []int{30, 540, 20},
		DemandList: []int{20, 30, 30, 10},
		CostTable: [][]int{
			{2, 3, 2, 4},
			{3, 2, 5, 1},
			{4, 3, 2, 6},
		},
	}

	jsonBlob, err := json.Marshal(inputTest)
	if err != nil {
		fmt.Println("Marshal error:", err)
	}
	fmt.Println("JSON:")
	fmt.Println(string(jsonBlob))
	return jsonBlob
}

func main() {
	params, err := ParseParams(testJSONData())

	if err != nil {
		fmt.Println("ParseParams error:", err)
		return
	}

	task := buildTaskFromParams(params)

	// Also here should be code which constructs this structure from parsed JSON
	task := Task{
		// It also needs some kind of stuct headerCell with value and potential maybe
		supplyList: []int{30, 540, 20},
		demandList: []int{20, 30, 30, 10},
		// Add tableCell struct that will contain cost and deliveryAmount and other fields
		// and use it as cells: [][]tableCell here
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

	// TODO: Validate parameters cost table dimensions and supply demand list dimensions

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
