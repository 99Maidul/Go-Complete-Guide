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

func (prod *Product) store() {
	file, err := os.OpenFile("./"+prod.id+".txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the product data to the file
	content := fmt.Sprintf("ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n", prod.id, prod.title, prod.description, prod.price)
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
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
	idInput = strings.TrimSuffix(idInput, "\n")
	idInput = strings.TrimSuffix(idInput, "\r")
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
	createdProudct.store()
}
