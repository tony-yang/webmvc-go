package tester

import (
  "strings"
  "testing"
)

func AssertStringEqual(t *testing.T, got, want string) {
  t.Helper()
	if strings.TrimSpace(got) != strings.TrimSpace(want) {
		t.Errorf("got string '%s', want '%s'", got, want)
	}
}

func AssertIntEqual(t *testing.T, got, want int) {
  t.Helper()
  if got != want {
    t.Errorf("got int '%d', want '%d'", got, want)
  }
}
