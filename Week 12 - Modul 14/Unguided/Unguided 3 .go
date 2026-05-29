package main

import "fmt"

type arrInt [1000]int

func insertionSort(arr *arrInt, n int) {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}

func main() {
	var data arrInt
	var n, x int

	// Membaca data sampai bilangan negatif
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[n] = x
		n++
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	if n > 1 {
		jarak := data[1] - data[0]
		tetap := true
		for i := 1; i < n-1; i++ {
			if data[i+1]-data[i] != jarak {
				tetap = false
				break
			}
		}

		if tetap {
			fmt.Printf("Data berjarak %d\n", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
