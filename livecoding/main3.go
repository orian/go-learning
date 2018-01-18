package livecoding

import "fmt"

func main() {
	fmt.Print("Podaj wiek: ")
	var wiek int
	fmt.Scan(&wiek)

	if wiek >= 21 {
		// i co teraz?
		fmt.Println("FULL ACCESS GRANTED")
	} else if wiek >= 18 {
		fmt.Println("RESTRICTED ACCESS GRANTED")
	} else {
		fmt.Println("ACCESS DENIED")
	}
	fmt.Println("the end.")
}
