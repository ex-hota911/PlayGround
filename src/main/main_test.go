package main

import "testing"

func TestMessage(t *testing.T) {
	a := Message()
	e := "Hello, world"
	if a != e {
		t.Errorf("got %v\nwanted %v", a, e)
	}
}
