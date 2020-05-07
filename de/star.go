package de

import (
	"fmt"
	"star"
)

// Sstar can
func Sstar(n int) string {
	return star.Sstar(n)
}

// Sstarf can
func Sstarf(n int) string {
	return fmt.Sprintf("%v (%v von %v)", Sstar(n), star.Min(n, 5), 5)
}
