package math

func Pow[T int | uint | uint8 | uint16 | uint32 | uint64](x, y T) T {
	if x == 0 {
		return 1
	}

	var acc = y
	for i := T(0); i < x-1; i++ {
		acc *= y
	}

	return acc
}

func Pow2[T int | uint | uint8 | uint16 | uint32 | uint64](x T) T {
	return Pow[T](x, 2)
}

func Pow3[T int | uint | uint8 | uint16 | uint32 | uint64](x T) T {
	return Pow[T](x, 3)
}

func Pow10[T int | uint | uint8 | uint16 | uint32 | uint64](x T) T {
	return Pow[T](x, 10)
}
