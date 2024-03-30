package main

import (
	"testing"
)

func TestMath(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado esperado %d, o que veio %d", 30, total)
	}
}
