package unit

var siPrefixes = map[string]float64{
	"Y": 1e24,
	"Z": 1e21,
	"E": 1e18,
	"P": 1e15,
	"T": 1e12,
	"G": 1e9,
	"M": 1e6,
	"k": 1e3,
	"h": 1e2,
	"da": 1e1,
	"d": 1e-1,
	"c": 1e-2,
	"m": 1e-3,
	"µ": 1e-6,
	"u": 1e-6,
	"n": 1e-9,
	"p": 1e-12,
	"f": 1e-15,
	"a": 1e-18,
	"z": 1e-21,
	"y": 1e-24,
}

var siUnits = map[string]struct{ unit theUnit }{
	"m/s2": {acceleration},
	"mol": {amountOfSubstance},
	"rad": {angle},
	"m2": {area},
	"bit/s": {datarate},
	"bit": {datasize},
	"s": {duration},
	"A": {electricCurrent},
	"F": {electricalCapacitance},
	"S": {electricalConductance},
	"H": {electricalInductance},
	"J": {energy},
	"m3/s": {flow},
	"N": {force},
	"Hz": {frequency},
	"lx": {illuminance},
	"m": {length},
	"lm": {luminousFlux},
	"g": {mass},
	"W": {power},
	"Pa": {pressure},
	"m/s": {speed},
	"K": {temperature},
	"V": {voltage},
	"m3": {volume},
}

func parseSymbol(symbol string) (float64, string) {
	longestPrefix := ""
	for prefix := range siPrefixes {
		if len(prefix) > len(longestPrefix) && len(symbol) >= len(prefix) && symbol[0:len(prefix)] == prefix {
			longestPrefix = prefix
		}
	}

	multiplier := 1.0
	unit := symbol

	if longestPrefix != "" {
		multiplier = siPrefixes[longestPrefix]
		unit = symbol[len(longestPrefix):]
	}

	return multiplier, unit
}
