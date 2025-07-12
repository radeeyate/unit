package unit

// ElectricCapacitance represents a SI unit of electric capacitance (in farad, F)
type ElectricCapacitance Unit

// ...
const (
	// SI
	Yoctofarad                 = Farad * 1e-24
	Zeptofarad                 = Farad * 1e-21
	Attofarad                  = Farad * 1e-18
	Femtofarad                 = Farad * 1e-15
	Picofarad                  = Farad * 1e-12
	Nanofarad                  = Farad * 1e-9
	Microfarad                 = Farad * 1e-6
	Millifarad                 = Farad * 1e-3
	Decifarad                  = Farad * 1e-2
	Centifarad                 = Farad * 1e-1
	Farad      ElectricCapacitance = 1e0
	Decafarad                  = Farad * 1e1
	Hectofarad                 = Farad * 1e2
	Kilofarad                  = Farad * 1e3
	Megafarad                  = Farad * 1e6
	Gigafarad                  = Farad * 1e9
	Terafarad                  = Farad * 1e12
	Petafarad                  = Farad * 1e15
	Exafarad                   = Farad * 1e18
	Zettafarad                 = Farad * 1e21
	Yottafarad                 = Farad * 1e24
)

// Yoctofarads returns the electrical capacitance in yF
func (c ElectricCapacitance) Yoctofarads() float64 {
	return float64(c / Yoctofarad)
}

// Zeptofarads returns the electrical capacitance in zF
func (c ElectricCapacitance) Zeptofarads() float64 {
	return float64(c / Zeptofarad)
}

// Attofarads returns the electrical capacitance in aF
func (c ElectricCapacitance) Attofarads() float64 {
	return float64(c / Attofarad)
}

// Femtofarads returns the electrical capacitance in fF
func (c ElectricCapacitance) Femtofarads() float64 {
	return float64(c / Femtofarad)
}

// Picofarads returns the electrical capacitance in pF
func (c ElectricCapacitance) Picofarads() float64 {
	return float64(c / Picofarad)
}

// Nanofarads returns the electrical capacitance in nF
func (c ElectricCapacitance) Nanofarads() float64 {
	return float64(c / Nanofarad)
}

// Microfarads returns the electrical capacitance in µF
func (c ElectricCapacitance) Microfarads() float64 {
	return float64(c / Microfarad)
}

// FromMillifarads return the value to be converted
func FromMillifarads(val float64) Value {
	return Value{val * float64(Millifarad), electricalCapacitance}
}

// toMillifarads return the converted value
func toMillifarads(value Value) (float64, error) {
	if value.unit != electricalCapacitance {
		return 0, ErrConversion
	}
	return ElectricCapacitance(value.val).Millifarads(), nil
}

// Millifarads returns the electrical capacitance in mF
func (c ElectricCapacitance) Millifarads() float64 {
	return float64(c / Millifarad)
}

// Decifarads returns the electrical capacitance in dF
func (c ElectricCapacitance) Decifarads() float64 {
	return float64(c / Decifarad)
}

// Centifarads returns the electrical capacitance in cF
func (c ElectricCapacitance) Centifarads() float64 {
	return float64(c / Centifarad)
}

// FromFarads return the value to be converted
func FromFarads(val float64) Value {
	return Value{val * float64(Farad), electricalCapacitance}
}

// toFarads return the converted value
func toFarads(value Value) (float64, error) {
	if value.unit != electricalCapacitance {
		return 0, ErrConversion
	}
	return ElectricCapacitance(value.val).Farads(), nil
}

// Farads returns the electrical capacitance in F
func (c ElectricCapacitance) Farads() float64 {
	return float64(c / Farad)
}

// Decafarads returns the electrical capacitance in daF
func (c ElectricCapacitance) Decafarads() float64 {
	return float64(c / Decafarad)
}

// Hectofarads returns the electrical capacitance in hF
func (c ElectricCapacitance) Hectofarads() float64 {
	return float64(c / Hectofarad)
}

// FromKilofarads return the value to be converted
func FromKilofarads(val float64) Value {
	return Value{val * float64(Kilofarad), electricalCapacitance}
}

// toKilofarads return the converted value
func toKilofarads(value Value) (float64, error) {
	if value.unit != electricalCapacitance {
		return 0, ErrConversion
	}
	return ElectricCapacitance(value.val).Kilofarads(), nil
}

// Kilofarads returns the electrical capacitance in kF
func (c ElectricCapacitance) Kilofarads() float64 {
	return float64(c / Kilofarad)
}

// Megafarads returns the electrical capacitance in MF
func (c ElectricCapacitance) Megafarads() float64 {
	return float64(c / Megafarad)
}

// Gigafarads returns the electrical capacitance in GF
func (c ElectricCapacitance) Gigafarads() float64 {
	return float64(c / Gigafarad)
}

// Terafarads returns the electrical capacitance in TF
func (c ElectricCapacitance) Terafarads() float64 {
	return float64(c / Terafarad)
}

// Petafarads returns the electrical capacitance in PF
func (c ElectricCapacitance) Petafarads() float64 {
	return float64(c / Petafarad)
}

// Exafarads returns the electrical capacitance in EF
func (c ElectricCapacitance) Exafarads() float64 {
	return float64(c / Exafarad)
}

// Zettafarads returns the electrical capacitance in ZF
func (c ElectricCapacitance) Zettafarads() float64 {
	return float64(c / Zettafarad)
}

// Yottafarads returns the electrical capacitance in YF
func (c ElectricCapacitance) Yottafarads() float64 {
	return float64(c / Yottafarad)
}
