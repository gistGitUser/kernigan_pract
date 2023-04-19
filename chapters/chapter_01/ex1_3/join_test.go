package main

import "testing"

var testArr = []string{"a", "b", "c", "d"}

func TestSimpleJoin(t *testing.T) {
	res := SimpleJoin(testArr, ",")
	resStd := StdJoin(testArr, ",")

	expOut := "a,b,c,d"

	if res != expOut {
		t.Fatal("unxpected out SimpleJoin get:" + res + " exp:" + expOut)
	}

	if resStd != expOut {
		t.Fatal("unxpected out StdJoin get:" + res + " exp:" + expOut)
	}

}

func BenchmarkSimpleJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleJoin([]string{"A man", "a plan", "a canal:", "Panama", "f", "et"}, ",")
	}
}

func BenchmarkStdJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdJoin([]string{"A man", "a plan", "a canal:", "Panama", "f", "et"}, ",")
	}
}
