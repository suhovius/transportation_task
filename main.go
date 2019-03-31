package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Remove this when frontend web server will be created
// that will handle these requests with such parameters
func testJSONData() []byte {
	inputTest := Params{
		SupplyList: []int{30, 40, 20},
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

	// TODO: Validate parameters cost table dimensions and supply demand list dimensions
	// TODO: Validate parameters

	task := buildTaskFromParams(params)

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

	// And might need to have added small floating point numbers here (0.001) to fix this degeneracy issue
	fmt.Println("Degeneracy Prevention: Add small amount to prevent degeneracy")
	task.preventDegeneracy()
	task.print()
	fmt.Println()

	fmt.Println("Base Plan: Calculated with 'North West Corner' method")
	task.northWestCorner()
	task.print()
	fmt.Printf("Delivery Cost: %f\n", task.deliveryCost())

	// Transport potentials method starts here. It should be loop

	// 1 Size distribution sum check

	// 2 Degeneracy check
	// http://cyclowiki.org/wiki/%D0%92%D1%8B%D1%80%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%BD%D0%BE%D1%81%D1%82%D1%8C_%D0%B2_%D1%82%D1%80%D0%B0%D0%BD%D1%81%D0%BF%D0%BE%D1%80%D1%82%D0%BD%D0%BE%D0%B9_%D0%B7%D0%B0%D0%B4%D0%B0%D1%87%D0%B5

	fmt.Print("\nDegeneracy Check:")
	if task.isDegenerate() {
		// TODO: Maybe return error here
		fmt.Println(" is Degenerate!")
		return
	}
	fmt.Println(" is not Degenerate")
	fmt.Println()

	// TODO: Round numners in api response generation and return int values there
	// https://yourbasic.org/golang/round-float-to-int/
}
