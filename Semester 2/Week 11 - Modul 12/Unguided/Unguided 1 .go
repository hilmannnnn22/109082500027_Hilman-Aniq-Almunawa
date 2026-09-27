package main

import "fmt"

type arrSuara [1000]int

func main() {
	var data arrSuara
	var suara [21]int
	var x int
	var jumlahMasuk int
	var jumlahSah int
	var idx int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		data[idx] = x
		idx++
		jumlahMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			jumlahSah++
		}
	}

	fmt.Println("Suara masuk:", jumlahMasuk)
	fmt.Println("Suara sah:", jumlahSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
