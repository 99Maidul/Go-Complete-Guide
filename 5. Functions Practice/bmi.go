package main

import (
	"github.com/99Maidul/BmiWithFunctions/info"
)

func main() {
	//Output Information
	info.PrintInfo()

	weight, height := getUserMetrics()

	//Calculate BMI
	bmi := calculateBMI(weight, height)

	//Output BMI
	printBMI(bmi)
}
func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
