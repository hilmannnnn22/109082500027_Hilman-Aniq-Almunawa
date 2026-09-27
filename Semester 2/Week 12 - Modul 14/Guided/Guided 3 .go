package main

import "fmt"

type arrInt [100]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]

		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var n, i int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan data:")
	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Println()
	fmt.Println("=== SEBELUM SORTING ===")
	for i = 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	insertionSort(&data, n)

	fmt.Println()
	fmt.Println("=== SETELAH SORTING DESCENDING ===")
	for i = 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
}
