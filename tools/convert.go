package tools

import "strconv"

func transFloat(a float64) string {
	return strconv.FormatFloat(a, 'f', 2, 64)
}

