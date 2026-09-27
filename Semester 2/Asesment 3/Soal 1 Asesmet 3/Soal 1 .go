package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func sorting(T *arrInt, n int) {
	var step, minIdx, j, sementara int
	for step = 0; step < n-1; step++ {
		minIdx = step
		for j = step + 1; j < n; j++ {
			if T[minIdx] > T[j] {
				minIdx = j
			}
		}
		sementara = T[minIdx]
		T[minIdx] = T[step]
		T[step] = sementara
	}
}

func median(T arrInt, n int) float64 {
	var nilaiTengah int = n / 2
	if n%2 != 0 {
		return float64(T[nilaiTengah])
	}
	return float64(T[nilaiTengah-1]+T[nilaiTengah]) / 2.0
}

func main() {
	var A arrInt
	var bilangan int
	var jumlah int = 0

	fmt.Println("Input data masukan : ")

	for {
		_, err := fmt.Scan(&bilangan)

		if err != nil || bilangan == -5313541 {
			break
		}

		if bilangan == 0 {
			sorting(&A, jumlah)
			fmt.Println("Median : ")
			fmt.Println(median(A, jumlah))
		} else {
			if jumlah < NMAX {
				A[jumlah] = bilangan
				jumlah++
			}
		}
	}
}
