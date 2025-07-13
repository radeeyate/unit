package unit

var fromMap = map[string]func(val float64) Value{
	"ac":      FromAcres,
	"ac-ft":   FromAcreFeet,
	"ac-in":   FromAcreInches,
	"ft3/s":   FromCubicFeetPerSecond,
	"gal":     FromUSLiquidGallons,
	"gal/h":   FromUSLiquidGallonsPerHour,
	"gal/min": FromUSLiquidGallonsPerMinute,
	"in-hg":   FromInchOfMercury,
	"psi":     FromPoundsPerSquareInch,
	"rpm":     FromRevolutionsPerMinute,
}

var toMap = map[string]func(value Value) (float64, error){
	"ac":      toAcres,
	"ac-ft":   toAcreFeet,
	"ac-in":   toAcreInches,
	"ft3/s":   toCubicFeetPerSecond,
	"gal":     toUSLiquidGallons,
	"gal/h":   toUSLiquidGallonsPerHour,
	"gal/min": toUSLiquidGallonsPerMinute,
	"in-hg":   toInchOfMercury,
	"psi":     toPoundsPerSquareInch,
	"rpm":     toRevolutionsPerMinute,
}
