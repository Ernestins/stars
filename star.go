package stars

import (
	"fmt"
)

// Min is more internal
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Sstar will
func Sstar(n int) string {
	result := ""
	for i := 0; i < Min(n, 5); i++ {
		result += fmt.Sprintf("⭐")
	}
	for i := Min(n, 5); i < 5; i++ {
		result += fmt.Sprintf("🌑")
	}
	return result
}

// SstarHTML will
func SstarHTML(n int, tag string) string {
	result := ""
	for i := 0; i < Min(n, 5); i++ {
		result += fmt.Sprintf("⭐")
	}
	for i := Min(n, 5); i < 5; i++ {
		result += fmt.Sprintf("🌑")
	}
	return fmt.Sprintf("<%v>%v</%v>", tag, result, tag)
}
