package stdconv

import "strconv"

func ParseInt(s string) int {
	number, _ := strconv.Atoi(s)

	return number
}
