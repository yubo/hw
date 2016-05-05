// Copyright 2016, yubo.
// All rights reserved.

package hw

import (
	"fmt"
	"testing"
)

func Example() {
	// example
	msg := Hw()
	fmt.Println(msg)
	// Output:
	// hello,world
}

func ExampleHw() {
	// example hw
	msg := Hw()
	fmt.Println(msg)
	// Output:
	// hello,world
}

func ExampleHwI() {
	// example hw interface
	msg := HwI(int(42))
	fmt.Println(msg)
	// Output:
	// hello,world int 42
}

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

func BenchmarkHw(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hw()
	}
}
