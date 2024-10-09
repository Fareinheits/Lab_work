package main

import "fmt"

func isPythagoreanTriplet(a int, b int, c int) bool {
	if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
		return true
	}
	return false
}

func main() {
	a := 3
	b := 4
	c := 5
	fmt.Println("Numbers:", a, b, c)
	fmt.Println("Is Pythagorean Triplet?", isPythagoreanTriplet(a, b, c))
}
