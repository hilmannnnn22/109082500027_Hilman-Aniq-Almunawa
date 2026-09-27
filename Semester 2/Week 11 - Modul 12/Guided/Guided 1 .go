package main

import "fmt"

func SequentialSearch(arrBuah [5]string, dataDicari string) int {
	var idx_found int = -1
	for i := 0; i < len(arrBuah); i++ {
		if arrBuah[i] == dataDicari {
			idx_found = i
			break
		}
	}
	return idx_found
}
func main() {
	var arrBuah [5]string

	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukkan data buah ke-%d: ", i)
		fmt.Scanln(&arrBuah[i])
	}

	var dataDicari string
	fmt.Print("Masukkan data buah yang ingin dicari: ")
	fmt.Scanln(&dataDicari)

	var index_data int
	index_data = SequentialSearch(arrBuah, dataDicari)

	if index_data > -1 {
		fmt.Printf("Data %s ditemukan pada index ke-%d\n", dataDicari, index_data)
	} else if index_data == -1 {
		fmt.Printf("Data %s tidak ditemukan\n", dataDicari)
	}
}
