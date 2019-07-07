package main

import (
	"github.com/AlecAivazis/survey"
	"os"
  "io/ioutil"
)

func main() {
  newArgs := os.Args[1:]

	response := ""
	prompt := &survey.Select{
		Message: newArgs[1],
		Options: newArgs[2:],
	}

	survey.AskOne(prompt, &response)
	//fmt.Println(response)
  ioutil.WriteFile(newArgs[0], []byte(response), 0644)
}
