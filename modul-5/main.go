package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func luas(a int, t int) int {
	luas := (a * t) / 2
	return luas
}

func main() {

	//Classic IF else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//IF without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//declare a variable inside IF condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//declare a variable inside IF condition
	//put && as "and" in IF condition statement
	if x := 10; x >= 0 && x <= 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	/*
		a := 10
		t := 10
		if a > t {
			fmt.Println("luas segitiganya adalah:", luas(a, t), ";Segitiga-nya Tinggi")
		} else if a == t {
			fmt.Println("luas segitiganya adalah:", luas(a, t), ";Segitiga-nya sama kaki")
		} else {
			fmt.Println("luas segitiganya adalah:", luas(a, t), ";Segitiga-nya Lebar")
		}
	*/

}
