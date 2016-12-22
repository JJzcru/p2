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
func PointerToBool(s *bool) bool {
	if s == nil {
		return false
	}
	return *(*bool)(unsafe.Pointer(s))
}
