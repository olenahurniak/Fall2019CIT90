// The mass m in pounds (lb) is equal to the mass m in kilograms (kg) divided by 0.45359237:
// m(lb) = m(kg) / 0.45359237

// The mass m in kilograms (kg) is equal to the mass m in pounds (lb) times 0.45359237:
// m(kg) = m(lb) × 0.45359237

package weightconv

// PToKg converts a Lb to Kg
func PToKg(p Pound) Kilogram { return Kilogram(p*0.45359237) }

// KgToP converts Kg to Lb
func KgToP(k Kilogram) Pound { return Pound(k/0.45359237) }