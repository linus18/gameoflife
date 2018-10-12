package main

import "testing"

func TestSomethoing(t *testing.T) {
	b := newBoard(12, 12)
	if b.getValue(1, 0) {
		t.Error("should be false")
	}
	b.setValue(10, 9)
	if !b.getValue(10, 9) {
		t.Error("should be true")
	}
}

func TestToString(t *testing.T) {
	b := newBoard(12, 12)
	b.setValue(0, 0)
	b.setValue(1, 1)
	b.setValue(1, 2)
	b.setValue(10, 10)
	println(b.toString())
}
