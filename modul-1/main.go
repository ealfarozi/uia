package main

import (
	"fmt"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Selamat belajar menggunakan Go!")
	fmt.Println("Waktu saat ini adalah: ", time.Now())
	//fmt.Println("Hello", "World!")

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
