package utils

import "math"

// RoundFloat arrotonda float64 a n decimali
func RoundFloat(val float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(val*multiplier) / multiplier
}

// PercentDifference calcola la differenza percentuale
func PercentDifference(actual, expected float64) float64 {
	if expected == 0 {
		return 0
	}
	return ((actual - expected) / expected) * 100
}
