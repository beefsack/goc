package goc

import "github.com/nsf/termbox-go"

type Text struct {
	Content []termbox.Cell
}

func (t *Text) Cells() ([]termbox.Cell, int) {
	width := 0
	lines := t.Lines()
	for _, l := range lines {
		if ll := len(l); ll > width {
			width = ll
		}
	}
	cells := []termbox.Cell{}
	for _, l := range lines {
		ll := len(l)
		cells = append(cells, l...)
		for i := 0; i < width-ll; i++ {
			cells = append(cells, termbox.Cell{Ch: 0})
		}
	}
	return cells, width
}

func (t *Text) Lines() [][]termbox.Cell {
	lines := [][]termbox.Cell{}
	line := []termbox.Cell{}
	for _, c := range t.Content {
		if c.Ch == '\n' {
			lines = append(lines, line)
			line = []termbox.Cell{}
		} else {
			line = append(line, c)
		}
	}
	return lines
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
