package livecoding

import "fmt"

func main() {
	fmt.Print("Podaj wiek: ")
	var wiek int
	fmt.Scan(&wiek)

	//
	fmt.Println("wiek:", wiek, ".")
	fmt.Printf("wiek: %d.", wiek)
	fmt.Printf(" a możę nie?")
}
