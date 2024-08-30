package resistorcolortrio

import (
	"math"
	"strconv"
)

const (
	kilo int = 1000
	mega     = kilo * 1000
	giga     = mega * 1000
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	mainValue := decode(colors[0])*10 + decode(colors[1])
	mainValue = mainValue * int(math.Pow(10.0, float64(decode(colors[2]))))
	resistor, unit := prefix(mainValue)
	return strconv.Itoa(resistor) + " " + unit + "ohms"
}

func decode(name string) int {
	colorMap := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	return colorMap[name]
}

func prefix(value int) (int, string) {
	if value == 0 {
		return 0, ""
	} else if value%giga == 0 {
		return value / giga, "giga"
	} else if value%mega == 0 {
		return value / mega, "mega"
	} else if value%kilo == 0 {
		return value / kilo, "kilo"
	}
	return value, ""
}
