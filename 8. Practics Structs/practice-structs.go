package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}
func main() {
	firstProduct := Product{
		"first-product,",
		"A book,",
		"A nice book,",
		10.99,
	}
	secondProduct := *NewProduct(
		"second-product,",
		"A carpet,",
		"A beautiful carpet,",
		99.99,
	)

	fmt.Println(firstProduct)
	fmt.Println(secondProduct)

}
