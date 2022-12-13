package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/99Maidul/bmi/info"
)

func main() {
	//Output Information
	fmt.Println(info.MainTittle)
	fmt.Println(info.MainTitleLine)

	//Prompt Input Information
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	//Save that info into variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	//Convert string to float
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	//Calculate BMI
	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", bmi)

}
