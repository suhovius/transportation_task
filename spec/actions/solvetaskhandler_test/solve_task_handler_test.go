package solvetaskhandler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/suhovius/transportation_task/app/actions/solvetaskhandler"
	"bitbucket.org/suhovius/transportation_task/app/forms/taskform"
	"github.com/go-resty/resty"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

// TODO: Cover all params validations cases
// TODO: Refactor all these support testing functions. Move them into another package

func assertJSONMatrixIntValues(t *testing.T, expectedMatrix [][]int64, jsonStr, path string) {
	for i, row := range gjson.Get(jsonStr, path).Array() {
		for j, cellVal := range row.Array() {
			assert.Equal(t, expectedMatrix[i][j], cellVal.Int())
		}
	}
}

func assertJSONArrayIntValues(t *testing.T, expectedArray []int64, jsonStr, path string) {
	for i, itemVal := range gjson.Get(jsonStr, path).Array() {
		assert.Equal(t, expectedArray[i], itemVal.Int())
	}
}

func assertJSONArrayBoolValues(t *testing.T, expectedArray []bool, jsonStr, path string) {
	for i, itemVal := range gjson.Get(jsonStr, path).Array() {
		assert.Equal(t, expectedArray[i], itemVal.Bool())
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

func newRequest(t *testing.T, ts *httptest.Server, taskParams *taskform.Params) (resp *resty.Response) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(taskParams).
		Post(ts.URL)

	require.Nil(t, err)

	return
}

func assertTaskCreateSuccess(t *testing.T, ts *httptest.Server, taskParams *taskform.Params) (result string) {
	// Act
	response := newRequest(t, ts, taskParams)

	// Assert
	assert.Equal(t, http.StatusOK, response.StatusCode())
	result = string(response.Body())

	return
}

func prepareValidParams() *taskform.Params {
	return &taskform.Params{
		SupplyList: []int{30, 40, 20},
		DemandList: []int{20, 30, 30, 10},
		CostTable: [][]int{
			{2, 3, 2, 4},
			{3, 2, 5, 1},
			{4, 3, 2, 6},
		},
	}
}

type ExpectedErrorResponse struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func assertErrorResponse(t *testing.T, eer *ExpectedErrorResponse, resp *resty.Response) {
	assert.Equal(t, resp.StatusCode(), eer.HTTPStatusCode)
	result := string(resp.Body())

	receivedErrorMessage := gjson.Get(result, "error_message").String()
	assert.Equal(t, eer.ErrorMessage, receivedErrorMessage)
}

func TestCreateTask(t *testing.T) {
	t.Log("with initialized server.")
	{
		logger := log.New()
		logger.Out = ioutil.Discard // disable server logging
		ts := httptest.NewServer(solvetaskhandler.New(logger))

		defer ts.Close()

		t.Log("\ttest:0\tcreates and processes task with valid params when there is more data in the input")
		{
			// Arrange
			validParams := taskform.Params{
				SupplyList: []int{30, 50, 75, 20},
				DemandList: []int{20, 40, 30, 10, 50, 25},
				CostTable: [][]int{
					{1, 2, 1, 4, 5, 2},
					{3, 3, 2, 1, 4, 3},
					{4, 2, 5, 9, 6, 2},
					{3, 1, 7, 3, 4, 6},
				},
			}

			// Act
			result := assertTaskCreateSuccess(t, ts, &validParams)

			// Assert
			exp := &TaskResponseExpectation{
				TotalDeliveryCost: 470,
				IsOptimalSolution: true,
				CostMatrix: [][]int64{
					{1, 2, 1, 4, 5, 2},
					{3, 3, 2, 1, 4, 3},
					{4, 2, 5, 9, 6, 2},
					{3, 1, 7, 3, 4, 6},
				},
				DeliveryAmountMatrix: [][]int64{
					{20, 0, 0, 0, 10, 0},
					{0, 0, 30, 0, 20, 0},
					{0, 40, 0, 0, 10, 25},
					{0, 0, 0, 10, 10, 0},
				},
				DemandAmounts:      []int64{20, 40, 30, 10, 50, 25},
				SupplyAmounts:      []int64{30, 50, 75, 20},
				SupplyIsFakeValues: []bool{false, false, false, false},
				DemandIsFakeValues: []bool{false, false, false, false, false, false},
				DemandPotentials:   []int64{1, 0, 0, 0, 5, 1},
				SupplyPotentials:   []int64{0, -1, 1, -1},
			}

			assertTaskCreateResponseBody(t, result, exp)
		}

		t.Log("\ttest:1\tcreates and processes task with valid params.")
		{
			// Arrange
			validParams := prepareValidParams()

			// Act
			result := assertTaskCreateSuccess(t, ts, validParams)

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
			validParams := taskform.Params{
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
			validParams := taskform.Params{
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

		t.Log("\ttest:4\treturns error when wrong request method have been used")
		{

			resp, err := resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(&taskform.Params{}).
				Get(ts.URL)

			require.Nil(t, err)

			// Assert
			assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode())
			result := string(resp.Body())

			receivedErrorMessage := gjson.Get(result, "error_message").String()
			assert.Equal(t, "Invalid request method", receivedErrorMessage)
		}

		t.Log("\ttest:5\treturns bad request error when broken json has been sent")
		{

			response, err := resty.R().
				SetBody(`this is not json`).
				SetHeader("Content-Type", "application/json").
				Post(ts.URL)

			require.Nil(t, err)

			// Assert
			eer := &ExpectedErrorResponse{
				HTTPStatusCode: http.StatusBadRequest,
				ErrorMessage:   "JSON Decoder: invalid character 'h' in literal true (expecting 'r')",
			}
			assertErrorResponse(t, eer, response)

		}
		t.Log("\ttest:6\treturns bad request error when cost table doesn't have any rows")
		{
			// Arrange
			taskParams := prepareValidParams()
			// Set empty cost table rows
			taskParams.CostTable = make([][]int, 0)

			// Act
			response := newRequest(t, ts, taskParams)

			// Assert
			eer := &ExpectedErrorResponse{
				HTTPStatusCode: http.StatusUnprocessableEntity,
				ErrorMessage:   "Params Validation Error: Cost Table should have at least one row",
			}
			assertErrorResponse(t, eer, response)
		}

		t.Log("\ttest:7\treturns bad request error when supply list is not equal to cost table rows count")
		{
			// Arrange
			taskParams := prepareValidParams()
			// Set wrong SupplyList elements size
			taskParams.SupplyList = []int{5, 3}

			// Act
			response := newRequest(t, ts, taskParams)

			// Assert
			eer := &ExpectedErrorResponse{
				HTTPStatusCode: http.StatusUnprocessableEntity,
				ErrorMessage: "Params Validation Error: Supply List size '2'" +
					" and Cost Table rows count '3' should be equal",
			}
			assertErrorResponse(t, eer, response)
		}

		t.Log("\ttest:8\treturns bad request error when demand list is not equal to cost table columns count")
		{
			// Arrange
			taskParams := prepareValidParams()
			// Set wrong DemandList elements size
			taskParams.DemandList = []int{9, 8, 42}

			// Act
			response := newRequest(t, ts, taskParams)

			// Assert
			eer := &ExpectedErrorResponse{
				HTTPStatusCode: http.StatusUnprocessableEntity,
				ErrorMessage: "Params Validation Error: Demand List size '3'" +
					" and Cost Table columns count '4' should be equal",
			}
			assertErrorResponse(t, eer, response)
		}

		t.Log("\ttest:9\treturns bad request error when cost table columns count is not correct at some row")
		{
			// Arrange
			taskParams := prepareValidParams()
			// Set wrong DemandList elements size
			taskParams.CostTable[2] = []int{7, 8}

			// Act
			response := newRequest(t, ts, taskParams)

			// Assert
			eer := &ExpectedErrorResponse{
				HTTPStatusCode: http.StatusUnprocessableEntity,
				ErrorMessage: "Params Validation Error: Cost Table row [2]" +
					" size '2' and Demand List size '4' should be equal",
			}
			assertErrorResponse(t, eer, response)
		}
	}
}
