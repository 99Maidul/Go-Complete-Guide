package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/99Maidul/BmiWithFunctions/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}
func getUserInput(promptText string) (value float64) {
	//Prompt Input Information
	fmt.Print(promptText)
	//Save that info into variables
	userinput, _ := reader.ReadString('\n')
	userinput = strings.Replace(userinput, "\n", "", -1)
	//Convert string to float
	value, _ = strconv.ParseFloat(userinput, 64)
	return
}
