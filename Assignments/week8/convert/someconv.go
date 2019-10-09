package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	OneKilogram Gramm = 1000
	OnePound    Gramm = 453.592
	OneOunce    Gramm = 28.35
)

func (k Kilogram) String() string { return fmt.Sprintf("%g°K", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g°P", p) }
