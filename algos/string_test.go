package algos

import "testing"

func TestBalancedBracket(t *testing.T) {
	str1 := "()"
	str2 := "()()"
	str3 := "[(){}]"
	str4 :="((({{{[[[]]]}}})))()(){}"

	if !BalancedBracket(str1) {
		t.Errorf("%s is balanced but program failed", str1)
	}
	if !BalancedBracket(str2) {
		t.Errorf("%s is balanced but program failed", str2)
	}
	if !BalancedBracket(str3) {
		t.Errorf("%s is balanced but program failed", str3)
	}
	if !BalancedBracket(str4) {
		t.Errorf("%s is balanced but program failed", str4)
	}
}

func TestBalancedBracket2(t *testing.T) {
	str1 := "{)"
	str2 := "(((((((())))"
	str3 := "{})((()"

	if BalancedBracket(str1) {
		t.Errorf("%s is unbalanced but program succeded", str1)
	}
	if BalancedBracket(str2) {
		t.Errorf("%s is unbalanced but program succeded", str2)
	}
	if BalancedBracket(str3) {
		t.Errorf("%s is unbalanced but program succeded", str3)
	}
}