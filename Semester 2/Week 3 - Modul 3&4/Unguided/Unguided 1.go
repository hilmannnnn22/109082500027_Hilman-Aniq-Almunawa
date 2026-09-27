package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int

	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
