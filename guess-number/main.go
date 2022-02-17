package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const finalMessage = "and press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playGame(firstNumber, secondNumber, subtraction, answer)
}

func playGame(firstNumber, secondNumber, subtraction, answer int) {
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
	fmt.Println("The answer is", answer)
}
