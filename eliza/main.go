package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var intro = doctor.Intro()

	fmt.Println(intro)
	for {
		fmt.Print("-> ")
		var userInput, _ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		}
		fmt.Println(doctor.Response(userInput))
	}
}
