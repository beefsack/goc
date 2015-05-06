package goc

import "github.com/nsf/termbox-go"

type Wrapped struct {
	Content Celler
}

func (w *Wrapped) Cells(x, y, width, height int) []termbox.Cell {
	return nil
}
