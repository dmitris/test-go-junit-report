package testfoo

import "testing"

func TestDoFoo(t *testing.T) {
	t.Parallel()
	if DoFoo("x") != "xx" {
		t.Error("Wrong result!")
	}
}

func TestDoFoo2(t *testing.T) {
	t.Parallel()
	if DoFoo("xyz") != "xyzxyz" {
		t.Error("Wrong result!")
	}
}
