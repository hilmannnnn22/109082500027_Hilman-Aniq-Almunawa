package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama, suara int
}
type tabPartai [NMAX]partai

func main() {
	var p tabPartai
	var n, inputNomor, posisiPartai int

	n = 0

	fmt.Println("Masukkan proses input suara : ")

	fmt.Scan(&inputNomor)
	for inputNomor != -1 {
		posisiPartai = posisi(p, n, inputNomor)
		if posisiPartai == -1 {
			p[n].nama = inputNomor
			p[n].suara = 1
			n++
		} else {
			p[posisiPartai].suara++
		}
		fmt.Scan(&inputNomor)
	}

	var i, j int
	var temp partai
	for i = 1; i < n; i++ {
		temp = p[i]
		j = i
		for j > 0 && p[j-1].suara < temp.suara {
			p[j] = p[j-1]
			j--
		}
		p[j] = temp
	}

	fmt.Println("\nHasil Perhitungan suara : ")
	for i = 0; i < n; i++ {
		fmt.Printf("%v(%v) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}

func posisi(t tabPartai, n int, nama int) int {
	var idx int
	for idx = 0; idx < n; idx++ {
		if t[idx].nama == nama {
			return idx
		}
	}
	return -1
}
