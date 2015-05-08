package goc

import "github.com/nsf/termbox-go"

type Padded struct {
	Content                  SizedCeller
	Top, Right, Bottom, Left int
}

func PaddedEven(content SizedCeller, padding int) *Padded {
	return &Padded{
		Content: content,
		Top:     padding,
		Right:   padding,
		Bottom:  padding,
		Left:    padding,
	}
}

func (p *Padded) SizedCells(x, y, width, height int) []termbox.Cell {
	cells := make([]termbox.Cell, width*height)
	if p.Content != nil {
		innerWidth := width - p.Left - p.Right
		innerHeight := height - p.Top - p.Bottom
		src := p.Content.SizedCells(x, y, innerWidth, innerHeight)
		CopyCells(src, cells, innerWidth, p.Left, p.Top, width)
	}
	return cells
}
