package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	var i int
	var ketemu bool = false

	i = 0
	for i < n && !ketemu {
		if T[i] == val {
			ketemu = true
		}
		i++
	}

	return ketemu
}

func inputSet(T *set, n *int) {
	var bil int

	*n = 0
	fmt.Scan(&bil)

	for *n < 2022 && !exist(*T, *n, bil) {
		T[*n] = bil
		*n++

		fmt.Scan(&bil)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	var i int

	*h = 0

	for i = 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			T3[*h] = T1[i]
			*h++
		}
	}
}

func printSet(T set, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)

	printSet(s3, n3)
}
