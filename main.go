package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println("sum is", add(a, b, c))

	avg := 0
	average(a, b, c, &avg)
	fmt.Println("average is", avg)

	min, max := min_max(a, b, c)
	fmt.Println("min number is", min)
	fmt.Println("max number is", max)

	radius := float64(b)
	fmt.Println("area of circle is",calculate_circle_area(radius))
}

func add(a, b, c int) int {
	return a + b + c
}
