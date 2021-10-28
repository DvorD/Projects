package awesomeProject

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter a:")
	var a float64
	fmt.Scanln(&a)

	z1 := (1 - 2 * math.Pow(math.Sin(a), 2)) / (1 + math.Sin(a + a))
	fmt.Println("z1 = ", z1)

	z2 := (1 - math.Tan(a)) / (1 + math.Tan(a))
	fmt.Println("z2 = ", z2)
}