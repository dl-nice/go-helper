package helper

import "strconv"

type I int

type I64 int64

type UI64 uint64

type ISwap interface {
	SwapToS() S
	SwapToF64() (F64, error)
}

func (i I) SwapToS() S {
	return S(strconv.Itoa(int(i)))
}

func (i I64) SwapToS() S {
	return S(strconv.FormatInt(int64(i), 10))
}

func (i UI64) SwapToS() S {
	return S(strconv.FormatUint(uint64(i), 10))
}

func (i I) SwapToF64() (F64, error) {
	f, err := strconv.ParseFloat(string(i.SwapToS()), 64)
	return F64(f), err
}

func (i I64) SwapToF64() (F64, error) {
	f, err := strconv.ParseFloat(string(i.SwapToS()), 64)
	return F64(f), err
}

func (i UI64) SwapToF64() (F64, error) {
	f, err := strconv.ParseFloat(string(i.SwapToS()), 64)
	return F64(f), err
}
