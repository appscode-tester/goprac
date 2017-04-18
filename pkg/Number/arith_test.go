package multiplication

import (
	"testing"
)

func TestNumberMul(t *testing.T) {
	var a, b int = 10, 20
	f1, _ := NewNumber(a)
	f2, _ := NewNumber(b)
	org, _ := f1.Mul(f2)
	exp, _ := NewNumber(a * b)
	if exp != org {
		t.Error("expeted: ", exp, " but found: ", org)
	}
	return
}

func TestNumberAdd(t *testing.T) {
	var a, b int = 10, 20
	f1, _ := NewNumber(a)
	f2, _ := NewNumber(b)
	org, _ := f1.Add(f2)
	exp, _ := NewNumber(a + b)
	if exp != org {
		t.Error("expeted: ", exp, " but found: ", org)
	}
	return
}
