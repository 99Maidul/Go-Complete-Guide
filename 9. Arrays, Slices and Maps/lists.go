package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func main() {
	var productNames [4]string = [4]string{"A book"}
	productNames[2] = "A carpet"
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
