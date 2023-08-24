package zzrand

import "testing"

func TestRNG_Intn(t *testing.T) {
	n := [10]int{}
	for i := 0; i < 1000000; i++ {
		n[Intn(10)]++
	}
	t.Log(n)
}
