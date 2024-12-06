package math

import "testing"

func TestfindMin(t *testing.T){
	test := findMin(1,2,3)
	res := 1
	if test != res {
		t.Errorf("function failed validation")
	}
}


func TestFindAverage(t *testing.T){
	test := findAverage(1,2,3)
	res := (1.0 + 2.0 + 3.0) / 3.0
	if test != res {
		t.Errorf("function failed validation")
	}
}

func TestFindFirstEquation(t *testing.T){
	test := findFirstEquation(2,9)
	res := -9 / 2
	if test != res {
		t.Errorf("function failed validation")
	}
}