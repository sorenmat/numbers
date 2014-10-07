package numbers

import "math"

/**
* This only works on positive numbers
 */
func Round(nr float64) int {
	flooredValue := float64(int(nr))

	x := nr - flooredValue

	if math.Abs(x) > 0.5 && nr > 0 {
		return int(nr) + 1
	} else if math.Abs(x) > 0.5 {

		return int(nr) - 1
	}

	return int(nr)

}
