package goc

import "github.com/nsf/termbox-go"

type SizedCeller interface {
	SizedCells(x, y, width, height int) []termbox.Cell
}

type Sized struct {
	Content Celler
}

func (s *Sized) SizedCells(x, y, width, height int) []termbox.Cell {
	cells := make([]termbox.Cell, width*height)
	if s.Content == nil {
		return cells
	}
	return cells
}
