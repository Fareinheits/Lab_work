package main

import "fmt"

func flatten(arr []interface{}) []int {
	var result []int
	for _, v := range arr {
		switch v := v.(type) {
		case int:
			result = append(result, v)
		case []interface{}:
			result = append(result, flatten(v)...)
		}
	}
	return result
}

func main() {
	nestedArr := []interface{}{1, 2, []interface{}{3, 4, []interface{}{5, 6}}, 7, []interface{}{8, 9}}
	fmt.Println("Nested Array:", nestedArr)
	fmt.Println("Flattened Array:", flatten(nestedArr))
}

// Task 1
