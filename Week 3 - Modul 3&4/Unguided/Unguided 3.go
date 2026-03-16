package main

import "fmt"

func cetakDeret(n int) {

	fmt.Print("Deret bilangan : ")

	for n != 1 {

		fmt.Print(n, " ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	fmt.Print(1)
}

func main() {

	var n int

	fmt.Print("Masukkan nilai awal : ")
	fmt.Scan(&n)

	cetakDeret(n)
}
