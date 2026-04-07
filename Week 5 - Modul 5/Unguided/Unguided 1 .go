package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 1)
}

func bintang(n int, i int) {
	if i > n {
		return
	} else {
		fmt.Println(cetakBintang(i))
		bintang(n, i+1)
	}
}

func cetakBintang(x int) string {
	if x == 0 {
		return ""
	} else {
		return "*" + cetakBintang(x-1)
	}
}