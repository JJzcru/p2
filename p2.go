package p2

import "unsafe"

//String transform a pointer to its string value
func String(s *string) string {
	if s == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(s))
}

//Bool transform a pointer to its bool value
func Bool(b *bool) bool {
	if b == nil {
		return false
	}
	return *(*bool)(unsafe.Pointer(b))
}

//Int transform a pointer to its int value
func Int(i *int) int {
	if i == nil {
		return 0
	}
	return *(*int)(unsafe.Pointer(i))
}
