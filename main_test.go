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
	b := newBoard(4, 8)
	b.setValue(1, 5)
	b.setValue(2, 4)
	b.setValue(2, 5)
	println(b.toString())
}

func TestNeighborCount(t *testing.T) {
	b := newBoard(4, 8)
	b.setValue(1, 5)
	b.setValue(2, 4)
	b.setValue(2, 5)
	if b.neighborCount(2, 5) != 2 {
		t.Error("Neighbor count should be 2")
	}
}

func TestNeighborCountLeftEdge(t *testing.T) {
	b := newBoard(4, 6)
	b.setValue(0, 0)
	//b.setValue(2, 0)
	if b.neighborCount(1, 0) != 1 {
		t.Error("Neighbor count should be 1")
	}
}

func TestNeighborCountRightEdge(t *testing.T) {
	b := newBoard(4, 6)
	b.setValue(0, 5)
	//b.setValue(2, 0)
	if b.neighborCount(1, 5) != 1 {
		t.Error("Neighbor count should be 1")
	}
}

func TestNeighborCountTop(t *testing.T) {
	b := newBoard(4, 6)
	b.setValue(0, 5)
	//b.setValue(2, 0)
	if b.neighborCount(0, 4) != 1 {
		t.Error("Neighbor count should be 1")
	}
}

func TestNeighborCountBottom(t *testing.T) {
	b := newBoard(4, 6)
	b.setValue(3, 5)
	//b.setValue(2, 0)
	if b.neighborCount(3, 4) != 1 {
		t.Error("Neighbor count should be 1")
	}
}

func TestNeighborCountCorner(t *testing.T) {
	b := newBoard(4, 6)
	b.setValue(0, 1)
	b.setValue(1, 0)
	if b.neighborCount(0, 0) != 2 {
		t.Error("Neighbor count should be 1")
	}
}
