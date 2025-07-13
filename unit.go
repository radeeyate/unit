package unit

import "errors"

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
	electricalCapacitance
	electricalConductance
	electricalInductance
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
	electricalResistance
)

var (
	ErrConversion = errors.New("wrong conversion unit")
	ErrNotFound   = errors.New("conversion unit not found")
)
