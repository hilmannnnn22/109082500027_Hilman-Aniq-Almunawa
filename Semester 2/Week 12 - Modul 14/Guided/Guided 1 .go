package main

import "fmt"

func SelectionSortArray(arr *[8]int) {
	var idx_min, i, j int

	for i = 0; i < len(arr)-1; i++ { // perulangan luar
		idx_min = i

		for j = i + 1; j < len(arr); j++ { // perulangan dalam
			if arr[j] < arr[idx_min] { // kondisi
				idx_min = j
			}
		}

		if idx_min != i { // swap
			arr[i], arr[idx_min] = arr[idx_min], arr[i]
		}
	}
}

func main() {
	var arrAngka [8]int
	var i int

	for i = 0; i < len(arrAngka); i++ {
		fmt.Printf("Masukkan data angka indeks ke-%d : ", i)
		fmt.Scan(&arrAngka[i])
	}

	fmt.Println()
	fmt.Println("=== SEBELUM SORTING ===")

	for i = 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}

	SelectionSortArray(&arrAngka)

	fmt.Println()
	fmt.Println("=== SETELAH SORTING ===")

	for i = 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}
}
