package unit

// ElectricInductance represents a SI unit of electric inductance (in henry, H)
type ElectricInductance Unit

// ...
const (
	// SI
	Yoctohenry                 = Henry * 1e-24
	Zeptohenry                 = Henry * 1e-21
	Attohenry                  = Henry * 1e-18
	Femtohenry                 = Henry * 1e-15
	Picohenry                  = Henry * 1e-12
	Nanohenry                  = Henry * 1e-9
	Microhenry                 = Henry * 1e-6
	Millihenry                 = Henry * 1e-3
	Decihenry                  = Henry * 1e-2
	Centihenry                 = Henry * 1e-1
	Henry      ElectricInductance = 1e0
	Decahenry                  = Henry * 1e1
	Hectohenry                 = Henry * 1e2
	Kilohenry                  = Henry * 1e3
	Megahenry                  = Henry * 1e6
	Gigahenry                  = Henry * 1e9
	Terahenry                  = Henry * 1e12
	Petahenry                  = Henry * 1e15
	Exahenry                   = Henry * 1e18
	Zettahenry                 = Henry * 1e21
	Yottahenry                 = Henry * 1e24
)

// Yoctohenries returns the electrical inductance in yH
func (c ElectricInductance) Yoctohenries() float64 {
	return float64(c / Yoctohenry)
}

// Zeptohenries returns the electrical inductance in zH
func (c ElectricInductance) Zeptohenries() float64 {
	return float64(c / Zeptohenry)
}

// Attohenries returns the electrical inductance in aH
func (c ElectricInductance) Attohenries() float64 {
	return float64(c / Attohenry)
}

// Femtohenries returns the electrical inductance in fH
func (c ElectricInductance) Femtohenries() float64 {
	return float64(c / Femtohenry)
}

// Picohenries returns the electrical inductance in pH
func (c ElectricInductance) Picohenries() float64 {
	return float64(c / Picohenry)
}

// Nanohenries returns the electrical inductance in nH
func (c ElectricInductance) Nanohenries() float64 {
	return float64(c / Nanohenry)
}

// Microhenries returns the electrical inductance in µH
func (c ElectricInductance) Microhenries() float64 {
	return float64(c / Microhenry)
}

// FromMillihenries return the value to be converted
func FromMillihenries(val float64) Value {
	return Value{val * float64(Millihenry), electricalInductance, 1}
}

// toMillihenries return the converted value
func toMillihenries(value Value) (float64, error) {
	if value.unit != electricalInductance {
		return 0, ErrConversion
	}
	return ElectricInductance(value.val).Millihenries(), nil
}

// Millihenries returns the electrical inductance in mH
func (c ElectricInductance) Millihenries() float64 {
	return float64(c / Millihenry)
}

// Decihenries returns the electrical inductance in dH
func (c ElectricInductance) Decihenries() float64 {
	return float64(c / Decihenry)
}

// Centihenries returns the electrical inductance in cH
func (c ElectricInductance) Centihenries() float64 {
	return float64(c / Centihenry)
}

// FromHenries return the value to be converted
func FromHenries(val float64) Value {
	return Value{val * float64(Henry), electricalInductance, 1}
}

// toHenries return the converted value
func toHenries(value Value) (float64, error) {
	if value.unit != electricalInductance {
		return 0, ErrConversion
	}
	return ElectricInductance(value.val).Henries(), nil
}

// Henries returns the electrical inductance in H
func (c ElectricInductance) Henries() float64 {
	return float64(c / Henry)
}

// Decahenries returns the electrical inductance in daH
func (c ElectricInductance) Decahenries() float64 {
	return float64(c / Decahenry)
}

// Hectohenries returns the electrical inductance in hH
func (c ElectricInductance) Hectohenries() float64 {
	return float64(c / Hectohenry)
}

// FromKilohenries return the value to be converted
func FromKilohenries(val float64) Value {
	return Value{val * float64(Kilohenry), electricalInductance, 1}
}

// toKilohenries return the converted value
func toKilohenries(value Value) (float64, error) {
	if value.unit != electricalInductance {
		return 0, ErrConversion
	}
	return ElectricInductance(value.val).Kilohenries(), nil
}

// Kilohenries returns the electrical inductance in kH
func (c ElectricInductance) Kilohenries() float64 {
	return float64(c / Kilohenry)
}

// Megahenries returns the electrical inductance in MH
func (c ElectricInductance) Megahenries() float64 {
	return float64(c / Megahenry)
}

// Gigahenries returns the electrical inductance in GH
func (c ElectricInductance) Gigahenries() float64 {
	return float64(c / Gigahenry)
}

// Terahenries returns the electrical inductance in TH
func (c ElectricInductance) Terahenries() float64 {
	return float64(c / Terahenry)
}

// Petahenries returns the electrical inductance in PH
func (c ElectricInductance) Petahenries() float64 {
	return float64(c / Petahenry)
}

// Exahenries returns the electrical inductance in EH
func (c ElectricInductance) Exahenries() float64 {
	return float64(c / Exahenry)
}

// Zettahenries returns the electrical inductance in ZH
func (c ElectricInductance) Zettahenries() float64 {
	return float64(c / Zettahenry)
}

// Yottahenries returns the electrical inductance in YH
func (c ElectricInductance) Yottahenries() float64 {
	return float64(c / Yottahenry)
}
