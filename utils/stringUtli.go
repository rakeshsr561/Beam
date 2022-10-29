package utils

import (
	"strings"
)

func AttachHashToLastDot(input string, hash string) string {
	lastInd := strings.LastIndex(input, ".")
	return input[:lastInd] + hash + input[lastInd+1:]
}
