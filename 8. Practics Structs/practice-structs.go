package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func getProduct() *Product {
	fmt.Println("Please enter your product data")
	fmt.Println("-----------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput, _ := readUserInput(reader, "Product ID:")
	titleInput, _ := readUserInput(reader, "Product Title:")
	descriptionInput, _ := readUserInput(reader, "Product Description:")
	priceInput, _ := readUserInput(reader, "Product Price:")
	priceInput = strings.TrimRight(priceInput, "\r")
	priceValue, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		fmt.Println("Error parsing product price:", err)
		return nil
	}

	product := Product{idInput, titleInput, descriptionInput, priceValue}
	return &product
}

func readUserInput(reader *bufio.Reader, promptText string) (string, error) {
	fmt.Print(promptText)
	userInput, err := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, err
}
func main() {
	createdProudct := getProduct()
	createdProudct.printData()
}
