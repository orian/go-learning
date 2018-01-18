// Mnożenie.
package main

import (
	"fmt"
	"math"
)

const pi = 3.14

func main() {
	r := 7.
	fmt.Println(2*pi*r)
	fmt.Println(2*math.Pi*r)
	area := 2*math.Pi*10
	fmt.Println(area)
}