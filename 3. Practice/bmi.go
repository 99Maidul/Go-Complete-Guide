package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//Output Information
	fmt.Println("BMI Calculator")
	fmt.Println("-------------")

	//Prompt Input Information
	fmt.Print("Enter your weight in kg: ")
	WeightInput, _ := reader.ReadString('\n')
	fmt.Print("Enter your height in m: ")
	HeightInput, _ := reader.ReadString('\n')

	//Save the input information
	WeightInput = strings.Replace(WeightInput, "\n", "", -1)
	HeightInput = strings.Replace(HeightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(WeightInput, 64)
	height, _ := strconv.ParseFloat(HeightInput, 64)

	fmt.Print(weight)
	fmt.Print(height)

}
