package gen

import (
	"regexp"
	"strings"

	"github.com/reiver/go-stringcase"
)

var nameReplacer = regexp.MustCompile(`[^\w]+`)

func toGoName(name string) string {
	return nameReplacer.ReplaceAllString(stringcase.ToPascalCase(name), "")
}

func notZeroByte(str string) string {
	if str != "0x00" {
		return str
	}
	return ""
}

func isReservedString(str string) bool {
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "reserved") {
		return true
	}

	if str == "res" {
		return true
	}

	return false
}
