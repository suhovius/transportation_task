package main

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

	task.print()
}
