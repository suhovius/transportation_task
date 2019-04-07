package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func assertJSONMatrixIntValues(t *testing.T, expectedMatrix [][]int64, jsonStr, path string) {
	for i, row := range gjson.Get(jsonStr, path).Array() {
		for j, cellVal := range row.Array() {
			assert.Equal(t, cellVal.Int(), expectedMatrix[i][j])
		}
	}
}

func assertJSONArrayIntValues(t *testing.T, expectedArray []int64, jsonStr, path string) {
	for i, itemVal := range gjson.Get(jsonStr, path).Array() {
		assert.Equal(t, itemVal.Int(), expectedArray[i])
	}
}

func assertJSONArrayBoolValues(t *testing.T, expectedArray []bool, jsonStr, path string) {
	for i, itemVal := range gjson.Get(jsonStr, path).Array() {
		assert.Equal(t, itemVal.Bool(), expectedArray[i])
	}
}

type TaskResponseExpectation struct {
	TotalDeliveryCost    int64
	IsOptimalSolution    bool
	CostMatrix           [][]int64
	DeliveryAmountMatrix [][]int64
	DemandAmounts        []int64
	SupplyAmounts        []int64
	SupplyIsFakeValues   []bool
	DemandIsFakeValues   []bool
	DemandPotentials     []int64
	SupplyPotentials     []int64
}

func assertTaskCreateResponseBody(t *testing.T, result string, exp *TaskResponseExpectation) {
	assert.Equal(
		t, gjson.Get(result, "total_delivery_cost").Int(), exp.TotalDeliveryCost,
	)

	assert.Equal(
		t, gjson.Get(result, "is_optimal_solution").Bool(), exp.IsOptimalSolution,
	)

	assertJSONMatrixIntValues(
		t, exp.CostMatrix, result, "table_cells.#.#.cost",
	)

	assertJSONMatrixIntValues(
		t, exp.DeliveryAmountMatrix, result, "table_cells.#.#.delivery_amount",
	)

	assertJSONArrayIntValues(
		t, exp.SupplyAmounts, result, "supply_list.#.amount",
	)

	assertJSONArrayIntValues(
		t, exp.SupplyPotentials, result, "supply_list.#.potential",
	)

	assertJSONArrayIntValues(
		t, exp.DemandAmounts, result, "demand_list.#.amount",
	)

	assertJSONArrayIntValues(
		t, exp.DemandPotentials, result, "demand_list.#.potential",
	)

	assertJSONArrayBoolValues(
		t, exp.SupplyIsFakeValues, result, "supply_list.#.is_fake",
	)

	assertJSONArrayBoolValues(
		t, exp.DemandIsFakeValues, result, "demand_list.#.is_fake",
	)
}

func assertTaskCreateSuccess(t *testing.T, ts *httptest.Server, taskParams *TaskParams) (result string) {
	// Act
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(taskParams).
		Post(ts.URL)

	require.Nil(t, err)

	// Assert
	assert.Equal(t, resp.StatusCode(), http.StatusOK)
	result = string(resp.Body())
	fmt.Printf("%v\n", result)

	return
}

func TestCreateTask(t *testing.T) {
	t.Log("with initialized server.")
	{
		ts := httptest.NewServer(&TaskSolvingHandler{})

		defer ts.Close()

		t.Log("\ttest:1\tcreates and processes task with valid params.")
		{
			// Arrange
			validParams := TaskParams{
				SupplyList: []int{30, 40, 20},
				DemandList: []int{20, 30, 30, 10},
				CostTable: [][]int{
					{2, 3, 2, 4},
					{3, 2, 5, 1},
					{4, 3, 2, 6},
				},
			}

			// Act
			result := assertTaskCreateSuccess(t, ts, &validParams)

			// Assert
			exp := &TaskResponseExpectation{
				TotalDeliveryCost: 170,
				IsOptimalSolution: true,
				CostMatrix: [][]int64{
					{2, 3, 2, 4},
					{3, 2, 5, 1},
					{4, 3, 2, 6},
				},
				DeliveryAmountMatrix: [][]int64{
					{20, 0, 10, 0},
					{0, 30, 0, 10},
					{0, 0, 20, 0},
				},
				DemandAmounts:      []int64{20, 30, 30, 10},
				SupplyAmounts:      []int64{30, 40, 20},
				SupplyIsFakeValues: []bool{false, false, false},
				DemandIsFakeValues: []bool{false, false, false, false},
				DemandPotentials:   []int64{2, 3, 2, 2},
				SupplyPotentials:   []int64{0, -1, 0},
			}

			assertTaskCreateResponseBody(t, result, exp)
		}

		t.Log("\ttest:2\tcreates and processes task with valid params when demand list is unbalanced")
		{
			// Arrange
			validParams := TaskParams{
				SupplyList: []int{30, 40, 20},
				DemandList: []int{20, 130, 30, 10},
				CostTable: [][]int{
					{2, 3, 2, 4},
					{3, 2, 5, 1},
					{4, 3, 2, 6},
				},
			}

			// Act
			result := assertTaskCreateSuccess(t, ts, &validParams)

			// Assert
			exp := &TaskResponseExpectation{
				TotalDeliveryCost: 170,
				IsOptimalSolution: true,
				CostMatrix: [][]int64{
					{2, 3, 2, 4},
					{3, 2, 5, 1},
					{4, 3, 2, 6},
					{0, 0, 0, 0},
				},
				DeliveryAmountMatrix: [][]int64{
					{20, 0, 10, 0},
					{0, 30, 0, 10},
					{0, 0, 20, 0},
					{0, 100, 0, 0},
				},
				DemandAmounts:      []int64{20, 130, 30, 10},
				SupplyAmounts:      []int64{30, 40, 20, 100},
				SupplyIsFakeValues: []bool{false, false, false, true},
				DemandIsFakeValues: []bool{false, false, false, false},
				DemandPotentials:   []int64{2, 3, 2, 2},
				SupplyPotentials:   []int64{0, -1, 0, -3},
			}

			assertTaskCreateResponseBody(t, result, exp)
		}

		t.Log("\ttest:3\tcreates and processes task with valid params when supply list is unbalanced")
		{
			// Arrange
			validParams := TaskParams{
				SupplyList: []int{30, 40, 520},
				DemandList: []int{20, 30, 30, 10},
				CostTable: [][]int{
					{2, 3, 2, 4},
					{3, 2, 5, 1},
					{4, 3, 2, 6},
				},
			}

			// Act
			result := assertTaskCreateSuccess(t, ts, &validParams)

			// Assert
			exp := &TaskResponseExpectation{
				TotalDeliveryCost: 220,
				IsOptimalSolution: true,
				CostMatrix: [][]int64{
					{2, 3, 2, 4, 0},
					{3, 2, 5, 1, 0},
					{4, 3, 2, 6, 0},
				},
				DeliveryAmountMatrix: [][]int64{
					{20, 0, 0, 0, 10},
					{0, 30, 0, 0, 10},
					{0, 0, 30, 10, 480},
				},
				DemandAmounts:      []int64{20, 30, 30, 10, 500},
				SupplyAmounts:      []int64{30, 40, 520},
				SupplyIsFakeValues: []bool{false, false, false},
				DemandIsFakeValues: []bool{false, false, false, false, true},
				DemandPotentials:   []int64{2, 0, 0, 0, 0},
				SupplyPotentials:   []int64{0, 0, 0},
			}

			assertTaskCreateResponseBody(t, result, exp)
		}
	}
}
