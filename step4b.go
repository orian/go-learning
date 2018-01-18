// Instrukcja warunkowa.
package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Podaj wiek: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Adult!")
	} else {
		fmt.Println("Not an adult ;(")
	}
}