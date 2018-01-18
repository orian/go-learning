package livecoding

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
)

func PoleKola(r float64) float64 {
	if r < 0 {
		panic("r < 0")
	}
	return r * r * math.Pi
}

var ErrRBelow0 = errors.New("r < 0")

// PoleKola2 oblicza pole kola. r musi być większe od 0.
func PoleKola2(r float64) (float64, error) {
	if r < 0 {
		return 0, ErrRBelow0
	}
	return r * r * math.Pi, nil
}

func PolePros(a, b float64) (float64, float64) {
	return a * b, 2 * (a + b)
}

func main() {
	//fmt.Println(PoleKola(11))
	//fmt.Println(PolePros(5, 6))
	a, err := PoleKola2(10)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Pole kola: ", a)
	}
}
