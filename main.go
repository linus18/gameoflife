package main

import (
	"bytes"
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

func main() {
	println("Hello there!")
}
