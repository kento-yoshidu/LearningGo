package main

import (
	"fmt"
	"os"
	"playground/calculator"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(3, 6))
}
