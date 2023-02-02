package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}
type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user.log.txt"}
	//do more work
	createLog(userLog)

	message := loggableString("[INFO] User Created!")
	// do more work
	createLog(message)
	outputValue(message)
	outputValue(userLog)
}

func createLog(data logger) {
	//more things to do
	data.log()
}

func outputValue(value interface{}) {
	val, ok := value.(string)
	fmt.Println(val, ok)
	fmt.Println(value)
}
