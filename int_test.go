package helper

import "testing"

func TestISwapToS(t *testing.T) {
	if I(1).SwapToS() != "1" {
		t.Errorf("SwapToS error,%v", I(1).SwapToS())
	}
	if I64(1).SwapToS() != "1" {
		t.Errorf("SwapToS error,%v", I64(1).SwapToS())
	}
	if UI64(1).SwapToS() != "1" {
		t.Errorf("SwapToS error,%v", UI64(1).SwapToS())
	}
}
func TestISwapToF64(t *testing.T) {
	if v, err := I(1).SwapToF64(); err != nil || v != F64(1) {
		t.Errorf("SwapToF64 error,%v,%v", v, err)
	}
	if v, err := I64(1).SwapToF64(); err != nil || v != F64(1) {
		t.Errorf("SwapToF64 error,%v,%v", v, err)
	}
	if v, err := UI64(1).SwapToF64(); err != nil || v != F64(1) {
		t.Errorf("SwapToF64 error,%v,%v", v, err)
	}
}
