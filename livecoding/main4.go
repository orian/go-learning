package livecoding

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	sekret := rand.Intn(100)
	var liczba int

	for liczba != sekret {
		fmt.Print("Podaj liczbę: ")
		fmt.Scan(&liczba)
		fmt.Println("Twoja liczba: ", liczba)

		if liczba < sekret {
			fmt.Println("Twoja liczba jest za mała.")
		} else {
			fmt.Println("Twoja liczba jest za duża.")
		}
	}
	fmt.Println("Super, zgadłeś poprawnie.")
}
