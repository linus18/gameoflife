package main

import (
	"bytes"
	"fmt"
	"time"
)

// Board type
type Board struct {
	grid [][]bool
	h, w int
}

func newBoard(h, w int) *Board {
	g := make([][]bool, h, w)
	for i := 0; i < h; i++ {
		g[i] = make([]bool, w)
	}
	return &Board{
		grid: g,
		h:    h,
		w:    w,
	}
}

func (b *Board) setValue(h int, w int) {
	b.grid[h][w] = true
}

func (b *Board) getValue(h int, w int) bool {
	if h < 0 || h >= b.h || w < 0 || w >= b.w {
		return false
	}
	return b.grid[h][w]
}

func (b *Board) toString() string {
	var buf bytes.Buffer
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			if b.grid[i][j] {
				buf.WriteByte('*')
			} else {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (b *Board) neighborCount(h, w int) int {
	cnt := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) && b.getValue(h+i, w+j) {
				cnt++
			}
		}
	}
	return cnt
}

func (b *Board) nextGen() {
	t := newBoard(b.h, b.w)
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			c := b.neighborCount(i, j)
			if c == 2 || c == 3 {
				if !(c == 2 && !b.getValue(i, j)) {
					t.setValue(i, j)
				}
			}
		}
	}
	b.grid = t.grid
}

// Well-known `Pulsar` pattern
func makePulsar() *Board {
	b := newBoard(17, 17)
	b.setValue(2, 4)
	b.setValue(2, 5)
	b.setValue(2, 6)
	b.setValue(2, 10)
	b.setValue(2, 11)
	b.setValue(2, 12)
	b.setValue(4, 2)
	b.setValue(4, 7)
	b.setValue(4, 9)
	b.setValue(4, 14)
	b.setValue(5, 2)
	b.setValue(5, 7)
	b.setValue(5, 9)
	b.setValue(5, 14)
	b.setValue(6, 2)
	b.setValue(6, 7)
	b.setValue(6, 9)
	b.setValue(6, 14)
	b.setValue(7, 4)
	b.setValue(7, 5)
	b.setValue(7, 6)
	b.setValue(7, 10)
	b.setValue(7, 11)
	b.setValue(7, 12)
	b.setValue(9, 4)
	b.setValue(9, 5)
	b.setValue(9, 6)
	b.setValue(9, 10)
	b.setValue(9, 11)
	b.setValue(9, 12)
	b.setValue(10, 2)
	b.setValue(10, 7)
	b.setValue(10, 9)
	b.setValue(10, 14)
	b.setValue(11, 2)
	b.setValue(11, 7)
	b.setValue(11, 9)
	b.setValue(11, 14)
	b.setValue(12, 2)
	b.setValue(12, 7)
	b.setValue(12, 9)
	b.setValue(12, 14)
	b.setValue(14, 4)
	b.setValue(14, 5)
	b.setValue(14, 6)
	b.setValue(14, 10)
	b.setValue(14, 11)
	b.setValue(14, 12)
	return b
}

func clear() {
	fmt.Println("\033[H\033[2J")
}

func main() {
	b := makePulsar()
	for i := 0; i < 60; i++ {
		clear()
		fmt.Println(b.toString())
		time.Sleep(1 * time.Second)
		b.nextGen()
	}
}
