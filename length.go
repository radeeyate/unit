package unit

// Length represents a SI unit of length (in meters, m)
type Length Unit

// ...
const (
	// SI
	Yoctometer              = Meter * 1e-24
	Zeptometer              = Meter * 1e-21
	Attometer               = Meter * 1e-18
	Femtometer              = Meter * 1e-15
	Picometer               = Meter * 1e-12
	Nanometer               = Meter * 1e-9
	Micrometer              = Meter * 1e-6
	Millimeter              = Meter * 1e-3
	Centimeter              = Meter * 1e-2
	Decimeter               = Meter * 1e-1
	Meter            Length = 1e0
	Decameter               = Meter * 1e1
	Hectometer              = Meter * 1e2
	Kilometer               = Meter * 1e3
	ScandinavianMile        = Meter * 1e4
	Megameter               = Meter * 1e6
	Gigameter               = Meter * 1e9
	Terameter               = Meter * 1e12
	Petameter               = Meter * 1e15
	Exameter                = Meter * 1e18
	Zettameter              = Meter * 1e21
	Yottameter              = Meter * 1e24
	Ronnameter              = Meter * 1e27
	Quettameter             = Meter * 1e30

	// US
	Inch    = Meter * 0.0254
	Hand    = Inch * 4
	Foot    = Inch * 12
	Yard    = Foot * 3
	Link    = Chain / 100
	Rod     = Yard * 5.5
	Chain   = Rod * 4
	Furlong = Chain * 10
	Mile    = Meter * 1609.344

	// US maritime
	Fathom       = Foot * 6
	Cable        = NauticalMile / 10
	NauticalMile = Meter * 1852

	// space
	LunarDistance    = Kilometer * 384400
	AstronomicalUnit = Meter * 149597870700
	LightYear        = Meter * 9460730472580800

	// chemistry
	Angstrom = Meter * 1e-10
)

// Yoctometers returns the length in ym
func (l Length) Yoctometers() float64 {
	return float64(l / Yoctometer)
}

// Zeptometers returns the length in zm
func (l Length) Zeptometers() float64 {
	return float64(l / Zeptometer)
}

// Attometers returns the length in am
func (l Length) Attometers() float64 {
	return float64(l / Attometer)
}

// Femtometers returns the length in fm
func (l Length) Femtometers() float64 {
	return float64(l / Femtometer)
}

// Picometers returns the length in pm
func (l Length) Picometers() float64 {
	return float64(l / Picometer)
}

// Nanometers returns the length in nm
func (l Length) Nanometers() float64 {
	return float64(l / Nanometer)
}

// Micrometers returns the length in microms
func (l Length) Micrometers() float64 {
	return float64(l / Micrometer)
}

// From Millimeters returns the value to be converted
func FromMillimeters(val float64) Value {
	return Value{val * float64(Millimeter), length, 1}
}

// toMillimeters return the converted value
func toMillimeters(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Millimeters(), nil
}

// Millimeters returns the length in mm
func (l Length) Millimeters() float64 {
	return float64(l / Millimeter)
}

// FromCentimeters return the value to be converted
func FromCentimeters(val float64) Value {
	return Value{val * float64(Centimeter), length, 1}
}

// toCentimeters return the converted value
func toCentimeters(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Centimeters(), nil
}

// Centimeters returns the length in cm
func (l Length) Centimeters() float64 {
	return float64(l / Centimeter)
}

// Decimeters returns the length in dm
func (l Length) Decimeters() float64 {
	return float64(l / Decimeter)
}

func FromMeters(val float64) Value {
	return Value{val, length, 1}
}

// toCentimeters return the converted value
func toMeters(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Meters(), nil
}

// Meters returns the length in m
func (l Length) Meters() float64 {
	return float64(l)
}

// Decameters returns the length in dam
func (l Length) Decameters() float64 {
	return float64(l / Decameter)
}

// Hectometers returns the length in hm
func (l Length) Hectometers() float64 {
	return float64(l / Hectometer)
}

// FromKilometers return the value to be converted
func FromKilometers(val float64) Value {
	return Value{val * float64(Kilometer), length, 1}
}

// toKilometers return the converted value
func toKilometers(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Kilometers(), nil
}

// Kilometers returns the length in km
func (l Length) Kilometers() float64 {
	return float64(l / Kilometer)
}

// ScandinavianMiles returns the length in in scandinavian miles (1 scandmile = 10 km)
func (l Length) ScandinavianMiles() float64 {
	return float64(l / ScandinavianMile)
}

// Megameters returns the length in in Mm
func (l Length) Megameters() float64 {
	return float64(l / Megameter)
}

// Gigameters returns the length in in Gm
func (l Length) Gigameters() float64 {
	return float64(l / Gigameter)
}

// Terameters returns the length in in Tm
func (l Length) Terameters() float64 {
	return float64(l / Terameter)
}

// Petameters returns the length in in Pm
func (l Length) Petameters() float64 {
	return float64(l / Petameter)
}

// Exameters returns the length in in Em
func (l Length) Exameters() float64 {
	return float64(l / Exameter)
}

// Zettameters returns the length in in Zm
func (l Length) Zettameters() float64 {
	return float64(l / Zettameter)
}

// Yottameters returns the length in in Ym
func (l Length) Yottameters() float64 {
	return float64(l / Yottameter)
}

// Ronnameters returns the length in in Rm
func (l Length) Ronnameters() float64 {
	return float64(l / Ronnameter)
}

// Quettameters returns the length in in Qm
func (l Length) Quettameters() float64 {
	return float64(l / Quettameter)
}

// FromInches return the value to be converted
func FromInches(val float64) Value {
	return Value{val * float64(Inch), length, 1}
}

// toInches return the converted value
func toInches(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Inches(), nil
}

// Inches returns the length in in
func (l Length) Inches() float64 {
	return float64(l / Inch)
}

// Hands returns the length in h
func (l Length) Hands() float64 {
	return float64(l / Hand)
}

// Feet returns the length in ft
func (l Length) Feet() float64 {
	return float64(l / Foot)
}

// Yards returns the length in yd
func (l Length) Yards() float64 {
	return float64(l / Yard)
}

// Rods returns the length in rod
func (l Length) Rods() float64 {
	return float64(l / Rod)
}

// Chains returns the length in ch
func (l Length) Chains() float64 {
	return float64(l / Chain)
}

// Links return the length in li
func (l Length) Links() float64 {
	return float64(l / Link)
}

// Furlongs returns the length in furlong
func (l Length) Furlongs() float64 {
	return float64(l / Furlong)
}

// FromMiles return the value to be converted
func FromMiles(val float64) Value {
	return Value{val * float64(Mile), length, 1}
}

// toMiles return the converted value
func toMiles(value Value) (float64, error) {
	if value.unit != length {
		return 0, ErrConversion
	}
	return Length(value.val).Miles(), nil
}

// Miles returns the length in mi
func (l Length) Miles() float64 {
	return float64(l / Mile)
}

// Fathoms returns the length in fathom
func (l Length) Fathoms() float64 {
	return float64(l / Fathom)
}

// Cables returns the length in cable
func (l Length) Cables() float64 {
	return float64(l / Cable)
}

// NauticalMiles returns the length in nm
func (l Length) NauticalMiles() float64 {
	return float64(l / NauticalMile)
}

// LunarDistances returns the length in ld
func (l Length) LunarDistances() float64 {
	return float64(l / LunarDistance)
}

// AstronomicalUnits returns the length in au
func (l Length) AstronomicalUnits() float64 {
	return float64(l / AstronomicalUnit)
}

// LightYears returns the length in ly
func (l Length) LightYears() float64 {
	return float64(l / LightYear)
}

// Angstroms returns the length in A
func (l Length) Angstroms() float64 {
	return float64(l / Angstrom)
}
