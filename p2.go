package p2

import "unsafe"

//PointerToString transform a pointer to its string value
func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(s))
}

//PointerToBool transform a pointer to its bool value
func PointerToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *(*bool)(unsafe.Pointer(b))
}

//PointerToInt transform a pointer to its int value
func PointerToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *(*int)(unsafe.Pointer(i))
}
