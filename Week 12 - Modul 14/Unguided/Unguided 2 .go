package main

import "fmt"

type arrInt [100000]int

func selectionSortGanjilGenap(arr *arrInt, m int) {
	var ganjil, genap arrInt
	var nGanjil, nGenap int

	for i := 0; i < m; i++ {
		if arr[i]%2 != 0 {
			ganjil[nGanjil] = arr[i]
			nGanjil++
		} else {
			genap[nGenap] = arr[i]
			nGenap++
		}
	}

	for i := 0; i < nGanjil-1; i++ {
		idx_min := i
		for j := i + 1; j < nGanjil; j++ {
			if ganjil[j] < ganjil[idx_min] {
				idx_min = j
			}
		}
		ganjil[i], ganjil[idx_min] = ganjil[idx_min], ganjil[i]
	}

	for i := 0; i < nGenap-1; i++ {
		idx_max := i
		for j := i + 1; j < nGenap; j++ {
			if genap[j] > genap[idx_max] {
				idx_max = j
			}
		}
		genap[i], genap[idx_max] = genap[idx_max], genap[i]
	}

	idx := 0
	for i := 0; i < nGanjil; i++ {
		arr[idx] = ganjil[i]
		idx++
	}
	for i := 0; i < nGenap; i++ {
		arr[idx] = genap[i]
		idx++
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

		selectionSortGanjilGenap(&data, m)

		for j := 0; j < m; j++ {
			fmt.Print(data[j], " ")
		}
		fmt.Println()
	}
}
