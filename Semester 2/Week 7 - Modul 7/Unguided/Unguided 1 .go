package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return (4.0 / 5.0) * c
}

func CelciusToFahrenheit(c suhu) suhu {
	return (9.0/5.0)*c + 32
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var c suhu

	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&c)

	fmt.Println("\n=== KONVERSI CELCIUS ===")
	fmt.Println(c, "celcius =", CelciusToReamur(c), "reamur")
	fmt.Println(c, "celcius =", CelciusToFahrenheit(c), "fahrenheit")
	fmt.Println(c, "celcius =", CelciusToKelvin(c), "kelvin")
}
