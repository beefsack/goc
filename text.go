package goc

import "github.com/nsf/termbox-go"

type Text struct {
	Content []termbox.Cell
}

func (t *Text) Cells(x, y, width, height int) []termbox.Cell {
	return t.Content
}
