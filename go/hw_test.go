// Copyright 2016, yubo.
// All rights reserved.

package hw

import (
	"testing"
)

func TestHw(t *testing.T) {
	out := Hw()
	if out != "hello,world" {
	}
	if got, want := Hw(), `hello,world`; got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHwI(t *testing.T) {
	tests := []struct {
		x    interface{}
		want string
	}{
		{bool(true), `hello,world bool true`},
		{float32(1.01), `hello,world float32 1.01`},
		{int(42), `hello,world int 42`},
	}
	for _, e := range tests {
		if got, want := HwI(e.x), e.want; got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}
