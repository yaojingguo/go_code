package algorithms

import "testing"

func TestIsPalindrome(t *testing.T) {
	// false
	if isPalindrome(-1) {
		t.Errorf("-1 is not  a palindrome")
	}
	if isPalindrome(12) {
		t.Errorf("12 is  not a palindrome")
	}

	// true
	if !isPalindrome(1) {
		t.Errorf("1 is  a palindrome")
	}
	if !isPalindrome(242) {
		t.Errorf("242 is  a palindrome")
	}
}

