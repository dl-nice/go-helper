package helper

import (
	"testing"
)

func TestSwapToI(t *testing.T) {
	if v, err := S("1").SwapToI(); err != nil || v != I(1) {
		t.Errorf("SwapToI error,%v,%v", v, err)
	}
}
func TestSwapToI64(t *testing.T) {
	if v, err := S("1").SwapToI64(); err != nil || v != I64(1) {
		t.Errorf("SwapToI64 error,%v,%v", v, err)
	}
}
func TestSwapToUI64(t *testing.T) {
	if v, err := S("1").SwapToUI64(); err != nil || v != UI64(1) {
		t.Errorf("SwapToUI64 error,%v,%v", v, err)
	}
}
func TestSwapToF32(t *testing.T) {
	if v, err := S("1").SwapToF32(); err != nil || v != F32(1) {
		t.Errorf("SwapToF32 error,%v,%v", v, err)
	}
}
func TestSwapToF64(t *testing.T) {
	if v, err := S("1").SwapToF64(); err != nil || v != F64(1) {
		t.Errorf("SwapToF64 error,%v,%v", v, err)
	}
}
