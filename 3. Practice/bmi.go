package main

import (
	"fmt"

	"github.com/99Maidul/bmi/info"
)

func main() {
	var height, weight float64
	//Output Information
	fmt.Println(info.MainTittle)
	fmt.Println(info.MainTitleLine)

	//Prompt Input Information
	fmt.Print(info.WeightPrompt)
	fmt.Scan(&weight)
	fmt.Print(info.HeightPrompt)
	fmt.Scan(&height)

	//Calculate BMI
	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", bmi)

}
