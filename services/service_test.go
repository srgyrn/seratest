package service

import "testing"

func TestIsDuck(t *testing.T) {
	t.Run("not duck", func(t *testing.T) {
		got := IsDuck("truck")
		if got != "nope! not a duck" {
			t.Error("expected something else here")
		}
	})

	t.Run("duck", func(t *testing.T) {
		got := IsDuck("duck")
		if got != "that is indeed a duck!" {
			t.Error("expected something else here")
		}
	})
}
