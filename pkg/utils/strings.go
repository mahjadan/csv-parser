package utils

import (
	"strings"
)

const ByteOrderMark = string('\uFEFF')

func ToLowerTrimSlice(columnSlice []string) []string {
	normalizedSlice := make([]string, len(columnSlice))
	for i, value := range columnSlice {
		if strings.HasPrefix(value, ByteOrderMark) {
			value = strings.TrimPrefix(value, ByteOrderMark)
		}
		normalizedSlice[i] = strings.TrimSpace(strings.ToLower(value))
	}
	return normalizedSlice
}
func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
