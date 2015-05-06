package goc

import "github.com/nsf/termbox-go"

type Padded struct {
	Content                  Celler
	Top, Right, Bottom, Left int
}

func PaddedEven(content Celler, padding int) *Padded {
	return &Padded{
		Content: content,
		Top:     padding,
		Right:   padding,
		Bottom:  padding,
		Left:    padding,
	}
}

func (p *Padded) Cells(x, y, width, height int) []termbox.Cell {
	cells := make([]termbox.Cell, width*height)
	if p.Content != nil {
		innerWidth := width - p.Left - p.Right
		CopyCells(
			p.Content.Cells(0, 0, innerWidth, height-p.Top-p.Bottom),
			cells, innerWidth, p.Left, p.Top, width,
		)
	}
	return cells
}
