package goc

import "github.com/nsf/termbox-go"

type Wrapped struct {
	Content  Celler
	WordWrap bool
}

func (w *Wrapped) SizedCells(x, y, width, height int) []termbox.Cell {
	cells := []termbox.Cell{}
	if w.Content == nil {
		return cells
	}
	return cells
}

func (w *Wrapped) WrappedCells(width int) []termbox.Cell {
	wrapped := []termbox.Cell{}
	return wrapped
}
