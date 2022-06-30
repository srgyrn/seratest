package greeter

import "testing"

func TestGreet_success(t *testing.T) {
	t.Log(Greet())
}

func TestGreet_fails(t *testing.T) {
	t.Error("oh no")
}
