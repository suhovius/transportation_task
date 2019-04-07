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

func TestCreateTask(t *testing.T) {
	t.Log("with initialized server.")
	{
		ts := httptest.NewServer(&TaskSolvingHandler{})

		defer ts.Close()

		t.Log("\ttest:1\tcreates and processes task with valid body.")
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
			resp, err := resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(validParams).
				Post(ts.URL)

			require.Nil(t, err)

			// Assert
			assert.Equal(t, resp.StatusCode(), http.StatusOK)
			result := string(resp.Body())
			fmt.Printf("%v\n", result)

			assert.Equal(
				t, gjson.Get(result, "total_delivery_cost").Int(), int64(170),
			)

			assert.Equal(
				t, gjson.Get(result, "is_optimal_solution").Bool(), true,
			)

			expectedCellCosts := [][]int64{
				{2, 3, 2, 4},
				{3, 2, 5, 1},
				{4, 3, 2, 6},
			}
			assertJSONMatrixIntValues(
				t, expectedCellCosts, result, "table_cells.#.#.cost",
			)

			expectedCellAmount := [][]int64{
				{20, 0, 10, 0},
				{0, 30, 0, 10},
				{0, 0, 20, 0},
			}
			assertJSONMatrixIntValues(
				t, expectedCellAmount, result, "table_cells.#.#.delivery_amount",
			)

			expectedSupplyAmount := []int64{30, 40, 20}
			assertJSONArrayIntValues(
				t, expectedSupplyAmount, result, "supply_list.#.amount",
			)

			expectedDemandAmount := []int64{20, 30, 30, 10}
			assertJSONArrayIntValues(
				t, expectedDemandAmount, result, "demand_list.#.amount",
			)

			expectedSupplyIsFake := []bool{false, false, false}
			assertJSONArrayBoolValues(
				t, expectedSupplyIsFake, result, "supplu_list.#.is_fake",
			)

			expectedDemandIsFake := []bool{false, false, false, false}
			assertJSONArrayBoolValues(
				t, expectedDemandIsFake, result, "demand_list.#.is_fake",
			)
		}
	}
}
