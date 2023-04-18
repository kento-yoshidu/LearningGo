package calculator

import "fmt"

var offset float64 = 4
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	fmt.Println("Multiply : ", Multiply(2, 3))
	fmt.Println("multiplay : ", multiply(2, 4))
	return a + b + offset
}
