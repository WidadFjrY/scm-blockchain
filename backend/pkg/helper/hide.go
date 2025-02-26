package helper

import "strings"

func HideAddress(address string) string {
	if len(address) < 6 {
		return address
	}

	visibleStart := address[:3]
	visibleEnd := address[len(address)-3:]
	hiddenPart := strings.Repeat("*", 3)

	return visibleStart + hiddenPart + visibleEnd
}
