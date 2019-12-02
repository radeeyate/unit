package unit

import (
	"testing"
)

func TestFlow(t *testing.T) {
	assertFloatEqual(t, 1e1, (10 * CubicMeterPerSecond).CubicMetersPerSecond())
	assertFloatEqual(t, 1e4, (10 * CubicMeterPerSecond).LitersPerSecond())
	assertFloatEqual(t, 600000, (10 * CubicMeterPerSecond).LitersPerMinute())
	assertFloatEqual(t, 3.6e+07, (10 * CubicMeterPerSecond).LitersPerHour())

	assertFloatEqual(t, 2641.7205235823217, (10 * CubicMeterPerSecond).USLiquidGallonsPerSecond())
	assertFloatEqual(t, 158503.2314149393, (10 * CubicMeterPerSecond).USLiquidGallonsPerMinute())
	assertFloatEqual(t, 9.510193884896357e+06, (10 * CubicMeterPerSecond).USLiquidGallonsPerHour())
	assertFloatEqual(t, 353.1466672148859, (10 * CubicMeterPerSecond).CubicFeetPerSecond())
}