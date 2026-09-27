package main

import "fmt"

func tambahValue(x int) {
	x = x + 10
	fmt.Println("Nilai x didalam prosedur tambahValue (pass by value) : ", x)
}

func tambahReference(x *int) {
	*x = *x + 10
	fmt.Println("Nilai x didalam prosedur tambahReference (pass by reference) : ", *x)
}

func main() {
	var y int = 5
	fmt.Println("Nilai awal : ", y)

	tambahValue(y)
	fmt.Println("Nilai y setelah pass by value : ", y)

	tambahReference(&y)
	fmt.Println("Nilai y setelah pass by reference : ", y)
}
