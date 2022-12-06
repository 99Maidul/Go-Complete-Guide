package info

import "fmt"

const mainTittle = "BMI Calculator"
const mainTitleLine = "-------------"
const WeightPrompt = "Enter your weight in (kg): "
const HeightPrompt = "Enter your height in (m): "

func PrintInfo() {
	fmt.Println(mainTittle)
	fmt.Println(mainTitleLine)
}
