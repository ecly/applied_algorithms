package hirsch

/*
 * Replicating the public tests from CodeJudge
 */

import (
	"testing"
)

func Hirschberg01Test(t *testing.T) {
	expected := "|||||"
	_, output := Hirschberg("abcde", "abcde")
	if output != expected {
		t.Errorf("Expected: %s, but was: %s", expected, output)
	}
}

func Hirschberg02Test(t *testing.T) {
	expected := "|||a||"
	_, output := Hirschberg("abcef", "abcdef")
	if output != expected {
		t.Errorf("Expected: %s, but was: %s", expected, output)
	}
}

func Hirschberg03Test(t *testing.T) {
	expected := "|||a||"
	_, output := Hirschberg("abcdef", "abcef")
	if output != expected {
		t.Errorf("Expected: %s, but was: %s", expected, output)
	}
}

func Hirschberg04Test(t *testing.T) {
	expected := "|b|||||||a|"
	_, output := Hirschberg("acdefkhijk", "abcdefghik")
	if output != expected {
		t.Errorf("Expected: %s, but was: %s", expected, output)
	}
}

func Hirschberg05Test(t *testing.T) {
	expected := "a||||bb"
	_, output := Hirschberg("xabcx", "abcabc")
	if output != expected {
		t.Errorf("Expected: %s, but was: %s", expected, output)
	}
}
