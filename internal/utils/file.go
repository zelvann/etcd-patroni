package utils

import (
	"strings"
)

func GetExtensions(filename string) string {
	return strings.Split(filename, ".")[1]
}
