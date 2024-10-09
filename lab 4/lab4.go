package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity int
}

func calculateTotalValue(products []Product) float64 {
	var totalValue float64
	for _, product := range products {
		totalValue += product.price * float64(product.quantity)
	}
	return totalValue
}

func main() {
	products := []Product{
		{"Продукт A", 10.99, 5},
		{"Продукт B", 9.99, 10},
		{"Продукт C", 12.99, 7},
	}
	fmt.Println("Продукти:")
	for _, product := range products {
		fmt.Printf("%s x %d = $%.2f\n", product.name, product.quantity, product.price*float64(product.quantity))
	}
	fmt.Println("Тотальна вартісь: $", calculateTotalValue(products))
}
