package numbers

import (
	"strconv"
	"strings"
)

func FloatDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}

	return 0
}
