package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)
}
