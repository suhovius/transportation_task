package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Remove this when frontend web server will be created
// that will handle these requests with such parameters
func testJSONData() []byte {
	inputTest := TaskParams{
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
	return jsonBlob
}

func main() {
	// ========= Parse Input JSON ==============================================
	// TODO: Parse json from server request parameters
	printLine()
	jsonBlob := testJSONData()
	fmt.Println("Received JSON:")
	fmt.Println(string(jsonBlob))

	params := TaskParams{}
	err := json.Unmarshal(jsonBlob, &params)

	if err != nil {
		fmt.Println("Parse request params error:", err)
		return
	}

	// ========= Parameters Validation =========================================
	// TODO: Validate parameters cost table dimensions and supply demand list dimensions
	// TODO: Validate parameters. At least one supply and at least one demand

	// ========= Create Task Struct ============================================

	task := (&TaskCreator{params: &params}).Perform()
	fmt.Printf("\nCreated Task UUID: %s\n", task.UUID)
	task.Print()

	// ========= Find the solution =============================================
	(&TaskSolver{task: &task}).Peform()
	// TODO: Round numners in api response generation and return int values there
	// https://yourbasic.org/golang/round-float-to-int/
}

func printLine() {
	fmt.Print("\n=========================================================\n\n")
}
