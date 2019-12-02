package unit

// Unit represents a unit
type Unit float64

type theUnit int

const (
	acceleration theUnit = iota + 1
	amountOfSubstance
	angle
	area
	datarate
	datasize
	duration
	electricCurrent
	electricalConductance
	electricalResistance
	energy
	flow
	force
	frequency
	illuminance
	length
	luminousFlux
	mass
	power
	pressure
	speed
	temperature
	voltage
	volume
)

var myMap = map[string]func(val float64) (theUnit, float64){
	"cm": func(val float64)(theUnit, float64){return length, val * float64(Centimeter)},
}
