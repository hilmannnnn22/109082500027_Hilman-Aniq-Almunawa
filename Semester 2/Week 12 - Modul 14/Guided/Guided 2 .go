package main

import "fmt"

type arrInt [100]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int

	for i = 1; i <= n-1; i++ {
		temp = T[i]
		j = i

		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
	}
}

func main() {
	var data arrInt
	var n, i int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
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
