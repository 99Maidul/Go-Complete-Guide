package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := getUserInput()
	storeData(input)
}

func getUserInput() string {
	fmt.Println("Please enter the text that should be stored")
	fmt.Println("Your input:")

	reader := bufio.NewReader(os.Stdin)

	enteredtext, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return enteredtext
}

func storeData(data string) {
	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Println("Storing data failed")
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Closing file failed")
		}
	}()

	file.WriteString(data)
	fmt.Println("Data stored successfully")

}
