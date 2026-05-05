package main

import "fmt"

type arrKelinci [1000]float64

func cariMinMax(arr arrKelinci, n int) (float64, float64) {
	var min float64 = arr[0]
	var max float64 = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max
}

func main() {
	var n int
	var data arrKelinci

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	min, max := cariMinMax(data, n)

	fmt.Println("Berat kelinci terkecil:", min)
	fmt.Println("Berat kelinci terbesar:", max)
}
