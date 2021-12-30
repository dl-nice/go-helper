package helper

import "strconv"

type F32 float32

type F64 float64

type FSwap interface {
	SwapToS() S
}

func (f F32) SwapToS() S {
	str := strconv.FormatFloat(float64(f), 'E', -1, 32)
	return S(str)
}

func (f F64) SwapToS() S {
	str := strconv.FormatFloat(float64(f), 'E', -1, 64)
	return S(str)
}
