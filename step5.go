// Gra.
package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	//rand.Seed(time.Now().Unix()) // włącz losowość`
	tajnaLiczba := rand.Intn(100)
	for {
		fmt.Printf("Wpisz liczbe: ")
		var liczba int
		fmt.Scanf("%d", &liczba)
		if liczba == tajnaLiczba {
			fmt.Println("Zgadłeś!")
			break
		} else if liczba > tajnaLiczba{
			fmt.Println("Za dużo, spróbuj ponownie.")
		} else {
			fmt.Println("Za mało, spróbuj ponownie.")
		}
	}
}