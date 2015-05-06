package goc

import "github.com/nsf/termbox-go"

type Celler interface {
	Cells(x, y, width, height int) []termbox.Cell
}

func SetCells(src []termbox.Cell, srcWidth, dstX, dstY int) {
	dst := termbox.CellBuffer()
	dstWidth, _ := termbox.Size()
	CopyCells(src, dst, srcWidth, dstX, dstY, dstWidth)
}

func CopyCells(src, dst []termbox.Cell, srcWidth, dstX, dstY, dstWidth int) {
	if src == nil {
		return
	}
	if dstX >= dstWidth {
		return
	}
	srcLen, dstLen := len(src), len(dst)
	srcPos, dstPos := 0, dstY*dstWidth+dstX
	if dstX == 0 && srcWidth == dstWidth {
		// Line widths match, just copy the whole thing at once.
		cpLen := srcLen
		if diff := dstPos + cpLen; diff < 0 {
			cpLen += diff
		}
		copy(dst[dstPos:], src[:cpLen])
		return
	}
	// Copy line by line.
	for srcPos < srcLen && dstPos < dstLen {
		take := srcWidth
		if diff := dstWidth - dstX + srcWidth; diff < 0 {
			take += diff
		}
		if diff := srcPos + srcLen - take; diff < 0 {
			take += diff
		}
		copy(dst[dstPos:], src[srcPos:srcPos+take])
		dstPos += dstWidth
		srcPos += srcWidth
	}
}

func StrToCells(input string, fg, bg termbox.Attribute) []termbox.Cell {
	return RunesToCells([]rune(input), fg, bg)
}

func RunesToCells(input []rune, fg, bg termbox.Attribute) []termbox.Cell {
	cells := make([]termbox.Cell, len(input))
	for i, r := range input {
		cells[i] = termbox.Cell{
			Ch: r,
			Fg: fg,
			Bg: bg,
		}
	}
	return cells
}

func Screen(c Celler) {
	width, height := termbox.Size()
	SetCells(c.Cells(0, 0, width, height), width, 0, 0)
}

type Sizer interface {
	Size() (width, height int)
}

type Widther interface {
	Width() int
}

func Width(c Celler) (width int, ok bool) {
	ok = true
	switch t := c.(type) {
	case Widther:
		width = t.Width()
	case Sizer:
		width, _ = t.Size()
	default:
		ok = false
	}
	return
}
