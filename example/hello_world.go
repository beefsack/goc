package main

import (
	"log"
	"time"

	"github.com/beefsack/goc"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatalf("failed to init termbox: %s", err)
	}
	defer termbox.Close()
	goc.CopyCells([]termbox.Cell{
		{Ch: 'E', Fg: termbox.ColorRed},
		{Ch: 'g'},
		{Ch: 'g'},
		{Ch: 0},
		{Ch: 0},
		{Ch: 'B'},
		{Ch: 'a'},
		{Ch: 'c'},
		{Ch: 'o'},
		{Ch: 'n'},
	}, 5, 2, 169, 44)
	if err := termbox.Flush(); err != nil {
		log.Fatalf("failed to flush, %s", err)
	}
	time.Sleep(5 * time.Second)
}
