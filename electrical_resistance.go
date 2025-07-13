package unit

// ElectricalResistance represents a SI derived unit of electrical resistance (in ohm, Ω)
type ElectricalResistance Unit

// ...
const (
	// SI
	Yoctoohm                 = Ohm * 1e-24
	Zeptoohm                 = Ohm * 1e-21
	Attoohm                  = Ohm * 1e-18
	Femtoohm                 = Ohm * 1e-15
	Picoohm                  = Ohm * 1e-12
	Nanoohm                  = Ohm * 1e-9
	Microohm                 = Ohm * 1e-6
	Milliohm                 = Ohm * 1e-3
	Centiohm                 = Ohm * 1e-2
	Deciohm                  = Ohm * 1e-1
	Ohm      ElectricalResistance = 1e0
	Decaohm                  = Ohm * 1e1
	Hectoohm                 = Ohm * 1e2
	Kiloohm                  = Ohm * 1e3
	Megaohm                  = Ohm * 1e6
	Gigaohm                  = Ohm * 1e9
	Teraohm                  = Ohm * 1e12
	Petaohm                  = Ohm * 1e15
	Exaohm                   = Ohm * 1e18
	Zettaohm                 = Ohm * 1e21
	Yottaohm                 = Ohm * 1e24
)

// Ohms returns the electrical resistance in Ω
func (e ElectricalResistance) Yoctoohms() float64 {
	return float64(e / Yoctoohm)
}

func (e ElectricalResistance) Zeptoohms() float64 {
	return float64(e / Zeptoohm)
}

func (e ElectricalResistance) Attoohms() float64 {
	return float64(e / Attoohm)
}

func (e ElectricalResistance) Femtoohms() float64 {
	return float64(e / Femtoohm)
}

func (e ElectricalResistance) Picoohms() float64 {
	return float64(e / Picoohm)
}

func (e ElectricalResistance) Nanoohms() float64 {
	return float64(e / Nanoohm)
}

func (e ElectricalResistance) Microohms() float64 {
	return float64(e / Microohm)
}

func (e ElectricalResistance) Milliohms() float64 {
	return float64(e / Milliohm)
}

func (e ElectricalResistance) Centiohms() float64 {
	return float64(e / Centiohm)
}

func (e ElectricalResistance) Deciohms() float64 {
	return float64(e / Deciohm)
}

func (e ElectricalResistance) Ohms() float64 {
	return float64(e)
}

func (e ElectricalResistance) Decaohms() float64 {
	return float64(e / Decaohm)
}

func (e ElectricalResistance) Hectoohms() float64 {
	return float64(e / Hectoohm)
}

func (e ElectricalResistance) Kiloohms() float64 {
	return float64(e / Kiloohm)
}

func (e ElectricalResistance) Megaohms() float64 {
	return float64(e / Megaohm)
}

func (e ElectricalResistance) Gigaohms() float64 {
	return float64(e / Gigaohm)
}

func (e ElectricalResistance) Teraohms() float64 {
	return float64(e / Teraohm)
}

func (e ElectricalResistance) Petaohms() float64 {
	return float64(e / Petaohm)
}

func (e ElectricalResistance) Exaohms() float64 {
	return float64(e / Exaohm)
}

func (e ElectricalResistance) Zettaohms() float64 {
	return float64(e / Zettaohm)
}

func (e ElectricalResistance) Yottaohms() float64 {
	return float64(e / Yottaohm)
}
