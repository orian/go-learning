// Wczytywanie danych
package main

import (
	"fmt"
)

func main() {
	var rokUrodzenia int
	fmt.Println("Podaj rok urodzenia: ")
	fmt.Scan(&rokUrodzenia)
	fmt.Println("Urodziles sie w: ", rokUrodzenia)
}