package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Print(WeightInput)
	fmt.Print(HeightInput)

}
