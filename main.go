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
	return jsonBlob
}

func main() {
	// ========= Parse Input JSON ==============================================
	// TODO: Parse json from server request parameters
	printLine()
	jsonBlob := testJSONData()
	fmt.Println("Received JSON:")
	fmt.Println(string(jsonBlob))
	params, err := ParseParams(jsonBlob)

	if err != nil {
		fmt.Println("ParseParams error:", err)
		return
	}
	printLine()

	// TODO: Refactor define serivce objects sturcts with methods for each step

	// TODO: Move this code into separate service object or smth / that later will be used at server at main method
	// ========= Parameters Validation =========================================
	// TODO: Validate parameters cost table dimensions and supply demand list dimensions
	// TODO: Validate parameters

	// ========= Create Task Struct ============================================
	task := buildTaskFromParams(params)
	fmt.Println("Initial State")
	task.Print()
	printLine()

	// ========= Initial Loop ==================================================
	var initialSteps []AlgorithmStep
	initialSteps = append(
		initialSteps,
		&Balancer{task: &task},
		&DegeneracyPreventer{task: &task},
		&NorthWestCornerSolutionFinder{task: &task},
	)

	initialPerformer := StepsSequencePerformer{task: &task, steps: &initialSteps}
	err = initialPerformer.Run()

	fmt.Printf("\nDelivery Cost: %d", roundToInt(task.deliveryCost()))

	// ========= Iterative Loop ================================================

	var iterativeSteps []AlgorithmStep
	iterativeSteps = append(
		iterativeSteps,
		&AmountDistributionChecker{task: &task},
		&DegeneracyChecker{task: &task},
		&PotentialsCalculator{task: &task},
		&OptimalSolutionChecker{task: &task},
		&CircuitBuilder{task: &task},
		&SupplyRedistributor{task: &task},
	)

	iterativePerformer := StepsSequencePerformer{
		task: &task, steps: &iterativeSteps,
	}

	err = iterativePerformer.Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// TODO: Check Cycles Count limit or finding time like 1 minute for example

	// TODO: Clear/Reset previous values from calculation

	// TODO: Round numners in api response generation and return int values there
	// https://yourbasic.org/golang/round-float-to-int/

	// Later each step could be started with step runner service object wrapper
	// which might perform loging and also might have config parameters regarding
	// what should be printed, to the log, and some others
}

func printLine() {
	fmt.Print("\n=========================================================\n\n")
}
