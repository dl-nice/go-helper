package helper

import "strconv"

type S string

func (s S) SwapToI() (I, error) {
	int, err := strconv.Atoi(string(s))
	return I(int), err
}

func (s S) SwapToI64() (I64, error) {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return I64(i), err
}

func (s S) SwapToUI64() (UI64, error) {
	i, err := strconv.ParseUint(string(s), 10, 64)
	return UI64(i), err
}

func (s S) SwapToF32() (F32, error) {
	f, err := strconv.ParseFloat(string(s), 32)
	return F32(f), err
}

func (s S) SwapToF64() (F64, error) {
	f, err := strconv.ParseFloat(string(s), 64)
	return F64(f), err
}
