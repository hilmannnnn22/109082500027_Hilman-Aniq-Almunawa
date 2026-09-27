package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri int = 0
	var kanan int = n - 1
	var tengah int

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
