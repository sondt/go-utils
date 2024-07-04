package go_utils

import "strconv"

// ToInt converts a string to int32
func ToInt(input string) int32 {
	number, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return 0
	}
	return int32(number)
}

// ToInt64 converts a string to int64
func ToInt64(input string) int64 {
	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0
	}
	return number
}

// ToFloat32 converts a string to float32
func ToFloat32(input string) float32 {
	number, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0
	}
	return float32(number)
}

// ToFloat64 converts a string to float64
func ToFloat64(input string) float64 {
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0
	}
	return number
}

// ToUint converts a string to uint32
func ToUint(input string) uint {
	number, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0
	}
	return uint(number)
}

// ToUint64 converts a string to uint64
func ToUint64(input string) uint64 {
	number, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0
	}
	return number
}
