package chart

import "math"

func Round(f float64) float64 {
	return math.Round(f*100) / 100
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}