package main

import "fmt"

type arrIkan [1000]float64

func main() {
	var x, y int
	var data arrIkan

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&data[i])
	}

	var totalWadah [1000]float64
	var jumlahWadah int
	var idx int = 0

	for idx < x {
		var total float64 = 0
		var count int = 0

		for count < y && idx < x {
			total += data[idx]
			idx++
			count++
		}

		totalWadah[jumlahWadah] = total
		jumlahWadah++
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(totalWadah[i], " ")
	}
	fmt.Println()

	var totalSemua float64
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalWadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)
	fmt.Println(rata)
}
