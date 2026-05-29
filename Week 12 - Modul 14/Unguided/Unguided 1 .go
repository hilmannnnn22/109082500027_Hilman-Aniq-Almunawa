package main

import "fmt"

type arrInt [100000]int

func selectionSort(arr *arrInt, m int) {
	var idx_min, temp int
	for i := 0; i < m-1; i++ {
		idx_min = i
		for j := i + 1; j < m; j++ {
			if arr[j] < arr[idx_min] {
				idx_min = j
			}
		}
		if idx_min != i {
			temp = arr[i]
			arr[i] = arr[idx_min]
			arr[idx_min] = temp
		}
	}
}

func main() {
	var n, m int
	var data arrInt

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
		}

		selectionSort(&data, m)

		for j := 0; j < m; j++ {
			fmt.Print(data[j], " ")
		}
		fmt.Println()
	}
}
