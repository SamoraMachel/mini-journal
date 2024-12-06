package lib

import "testing"

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q wanted %q", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one: %q", got)
	}
}