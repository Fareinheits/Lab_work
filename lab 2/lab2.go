package main

import "fmt"

func findCommonElements(arr1 []int, arr2 []int) []int {
	var result []int
	for _, v := range arr1 {
		for _, w := range arr2 {
			if v == w {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	fmt.Println("Массив 1:", arr1)
	fmt.Println("Массив 2:", arr2)
	fmt.Println("Звичайні елементи:", findCommonElements(arr1, arr2))
}
