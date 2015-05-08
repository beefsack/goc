package goc

import "github.com/nsf/termbox-go"

type HeightCeller interface {
	Celler
	Height(width int) int
}

type VertScroll struct {
	Content HeightCeller
}

func (vs *VertScroll) Cells(x, y, width, height int) []termbox.Cell {
	return nil
}
