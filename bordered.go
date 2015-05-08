package goc

import (
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	BorderT = iota
	BorderTR
	BorderR
	BorderBR
	BorderB
	BorderBL
	BorderL
	BorderTL
)

type BorderRunes map[int]rune

var defaultBorderRunes = BorderRunes{
	BorderT:  '─',
	BorderTR: '┐',
	BorderR:  '│',
	BorderBR: '┘',
	BorderB:  '─',
	BorderBL: '└',
	BorderL:  '│',
	BorderTL: '┌',
}

type Bordered struct {
	Content                              Celler
	Label                                string
	Runes                                BorderRunes
	BorderFg, BorderBg, LabelFg, LabelBg termbox.Attribute
}

func (b *Bordered) borderRune(dir int) rune {
	if b.Runes != nil {
		if r, ok := b.Runes[dir]; ok {
			return r
		}
	}
	return defaultBorderRunes[dir]
}

func (b *Bordered) SizedCells(x, y, width, height int) []termbox.Cell {
	if height <= 0 || width <= 0 {
		return []termbox.Cell{}
	}
	runes := b.Runes
	if runes[0] == 0 {
		runes = defaultBorderRunes
	}
	cells := make([]termbox.Cell, width*height)
	if b.Content != nil {
		src, srcWidth := b.Content.Cells()
		CopyCells(src, cells, srcWidth, 1, 1, width)
	}
	cells[0] = termbox.Cell{
		Ch: b.borderRune(BorderTL),
		Fg: b.BorderFg,
		Bg: b.BorderBg,
	}
	cells[width-1] = termbox.Cell{
		Ch: b.borderRune(BorderTR),
		Fg: b.BorderFg,
		Bg: b.BorderBg,
	}
	CopyCells(StrToCells(
		strings.Repeat(string(b.borderRune(BorderT)), width-2),
		b.BorderFg,
		b.BorderBg,
	), cells, width, 1, 0, width)
	for i := 1; i < height-1; i++ {
		lineStart := width * i
		cells[lineStart] = termbox.Cell{
			Ch: b.borderRune(BorderL),
			Fg: b.BorderFg,
			Bg: b.BorderBg,
		}
		cells[lineStart+width-1] = termbox.Cell{
			Ch: b.borderRune(BorderR),
			Fg: b.BorderFg,
			Bg: b.BorderBg,
		}
	}
	if height > 1 {
		lineStart := width * (height - 1)
		cells[lineStart] = termbox.Cell{
			Ch: b.borderRune(BorderBL),
			Fg: b.BorderFg,
			Bg: b.BorderBg,
		}
		cells[lineStart+width-1] = termbox.Cell{
			Ch: b.borderRune(BorderBR),
			Fg: b.BorderFg,
			Bg: b.BorderBg,
		}
		CopyCells(StrToCells(
			strings.Repeat(string(b.borderRune(BorderB)), width-2),
			b.BorderFg,
			b.BorderBg,
		), cells, width, 1, height-1, width)
	}
	if b.Label != "" {
		CopyCells(StrToCells(b.Label, b.LabelFg, b.LabelBg), cells,
			width-2, 1, 0, width-1)
	}
	return cells
}
