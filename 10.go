package awesomeProject

import (
"fmt"
"math"
)

func main() {
	fmt.Println("Enter a:")
	var a float64
	fmt.Scanln(&a)

	b1 := (math.Pi / 2) + 3*a
	b2 := 3*a - math.Pi
	z1 := math.Sin(b1) / (1 - math.Sin(b2))
	fmt.Println("z1 = ", z1)

	c := (5 / 4 * math.Pi) + (3 / 2 * a)
	z2 := math.Cos(c) / math.Sin(c)
	fmt.Println("z2 = ", z2)
}