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

	// TODO: Move this code into separate service object or smth / that later will be used at server at main method
	// ========= Parameters Validation =========================================
	// TODO: Validate parameters cost table dimensions and supply demand list dimensions
	// TODO: Validate parameters

	// ========= Create Task Struct ============================================
	task := buildTaskFromParams(params)
	fmt.Println("Initial State")
	task.Print()
	printLine()

	// ========= Perform Balancing =============================================
	kind := task.performBalancing()
	switch kind {
	case "nothing":
		fmt.Println("Balancing: Task is already balanced. Skip balancing")
	case "fake_demand":
		fmt.Println("Balancing: Add fake demand")
	case "fake_supply":
		fmt.Println("Balancing: Add fake supply")
	}
	task.Print()
	printLine()

	// ========= Degeneracy Prevention =========================================
	fmt.Println("Degeneracy Prevention: Add small amount to prevent degeneracy")
	task.preventDegeneracy()
	// TODO: Make Print configurable with ability to print (or Print and PrintAccurate)
	// read float values in the table to show these small disturbances
	// that were added at this step
	task.Print()
	printLine()

	fmt.Println("Base Plan: Calculated with 'North West Corner' method")
	task.northWestCorner()
	task.Print()
	fmt.Printf("\nDelivery Cost: %d", roundToInt(task.deliveryCost()))
	printLine()

	// ========= Amount distribution sum check =================================
	fmt.Print("Amount distribution sum check. ")
	err = task.amountDistributionCheck()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sum matched")
	printLine()

	// ========= Degeneracy Check ==============================================
	fmt.Print("Degeneracy Check:")
	if task.isDegenerate() {
		// TODO: Maybe return error here
		fmt.Print(" is Degenerate!")
		return
	}
	fmt.Print(" is not Degenerate")
	printLine()

	// ========= Potentials Calculation ========================================
	fmt.Println("Potentials Calculation")
	task.calculatePotentials()
	task.Print()
	printLine()

	// ========= Optimal Solution Check ========================================
	fmt.Println("Optimal Solution Check")
	isOptimal := task.optimalSolutionCheck()
	task.Print()
	if isOptimal {
		fmt.Println("is optimal")
	} else {
		fmt.Println("is not optimal")
		i := task.MinDeltaCell.i
		j := task.MinDeltaCell.j
		fmt.Printf(
			"Min Negative Delta Cell: D[%d][%d]= %d\n", i, j,
			roundToInt(task.tableCells[i][j].delta),
		)
	}
	printLine()

	// ========= Build Circuit ===================================================
	fmt.Println("Build Circuit")
	// TODO: Refactor define serivce objects sturcts with methods for each step
	// like this CircuitBuilder{taskPtr: &task}.run()
	// TODO: Google Golang service objects!!!

	// This is the idea of the processing conveyor!!!
	// var steps []AlgorithmStep
	// Needs to have task to be set there also
	// steps = append(steps, &CircuitBuilder{task: &task}, &SupplyRedistributor{task: &task}, &OtherStep{task: &task})
	// for _, step := range steps {
	//  Maybe we can add step.Init(task: &task) or smth to let it work properly
	// or just provide task as the argument of the Perform method
	// 	step.Perform()
	//  here also migth happen logging inside another service object wrapper
	// }

	// var s []*MyStruct

	// for i := 0; i < b.N; i++ {
	//     s = append(s, &MyStruct{})
	// }

	(&CircuitBuilder{task: &task}).Perform()
	// taskPrinter{taskPtr: &task}.perform()

	// TODO: Google Golang service objects!!!
	// TODO: Maybe Goroutines can be used for Circuit Building

	// TODO: Round numners in api response generation and return int values there
	// https://yourbasic.org/golang/round-float-to-int/

	// Each step should be implement interface with method Run() or Perform()
	// and might return error
	// Later each step could be started with step runner service object wrapper
	// which might perform loging and alos might have config parameters regarding
	// what should be printed, to the log, and some others
	// var steps *[]AlgorithmStep = [CircuitBuilder, SupplyRdistributor]
}

func printLine() {
	fmt.Print("\n=====================================================================================\n\n")
}
