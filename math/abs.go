package math

func Abs[T int | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
