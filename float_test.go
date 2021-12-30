package helper

import (
	"testing"
)

func TestFSwapToS(t *testing.T) {
	if F32(3.1415926535).SwapToS() != "3.1415927E+00" {
		t.Errorf("SwapToS error,%v", F64(3.1415926535).SwapToS())
	}
	if F64(3.1415926535).SwapToS() != "3.1415926535E+00" {
		t.Errorf("SwapToS error,%v", F64(3.1415926535).SwapToS())
	}
}
