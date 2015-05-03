package goc

import "github.com/nsf/termbox-go"

func CopyCells(src []termbox.Cell, srcWidth, srcHeight, dstX, dstY int) {
	if src == nil {
		return
	}
	srcLen := len(src)
	dstWidth, _ := termbox.Size()
	if dstX >= dstWidth {
		return
	}
	dst := termbox.CellBuffer()
	dstLen := len(dst)
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
		copy(dst[dstPos:], src[srcPos:srcPos+take])
		dstPos += dstWidth
		srcPos += srcWidth
	}
}
