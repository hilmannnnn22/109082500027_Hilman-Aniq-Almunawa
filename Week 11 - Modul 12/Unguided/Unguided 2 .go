package main

import "fmt"

type arrSuara [21]int

func main() {
	var suara arrSuara
	var x int
	var suaraMasuk int
	var suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	var ketua int = 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	var wakil int

	if ketua == 1 {
		wakil = 2
	} else {
		wakil = 1
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
