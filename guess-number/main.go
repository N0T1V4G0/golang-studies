package main

import (
	"bufio"
	"fmt"
	"os"
)

const finalMessage = "and press ENTER when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("GUESS THE NUMBER GAME")
	fmt.Println("_____________________")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", finalMessage)
	reader.ReadString('\n')
	fmt.Println("Multiply the number by", firstNumber, finalMessage)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", secondNumber, finalMessage)
	reader.ReadString('\n')
	fmt.Println("Divide the result by the number you thought of", finalMessage)
	reader.ReadString('\n')
	fmt.Println("Now subtract the result by", subtraction, finalMessage)
	reader.ReadString('\n')
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
