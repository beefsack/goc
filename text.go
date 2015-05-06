package goc

import "github.com/nsf/termbox-go"

type Text struct {
	Content []termbox.Cell
}

func (t *Text) Cells(x, y, width, height int) []termbox.Cell {
	cells := make([]termbox.Cell, width*height)
	if t.Content == nil {
		return cells
	}
	curX, curY := 0, 0
	cellPos := 0
	for _, c := range t.Content {
		if c.Ch == '\n' {
			if diff := x + width - curX; diff > 0 {
				cellPos += diff
			}
			curX = 0
			curY++
			if curY >= y+height {
				break
			}
		} else if curX >= x && curX < x+width && curY >= y && curY < y+height {
			cells[cellPos] = c
			cellPos++
			curX++
		}
	}
	return cells
}

func (t *Text) Size() (width, height int) {
	if t.Content == nil {
		return
	}
	line := 0
	for _, c := range t.Content {
		if c.Ch == '\n' {
			height++
			if line > width {
				width = line
			}
		}
	}
	return
}
