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
	goc.Screen(&goc.Bordered{
		Content: &goc.Text{
			Content: []termbox.Cell{
				{Ch: 'E', Fg: termbox.ColorRed},
				{Ch: 'g'},
				{Ch: 'g'},
			},
		},
		Label:    "Check out the border!",
		BorderFg: termbox.ColorGreen,
		LabelFg:  termbox.ColorBlue | termbox.AttrUnderline | termbox.AttrBold,
	})
	if err := termbox.Flush(); err != nil {
		log.Fatalf("failed to flush, %s", err)
	}
	time.Sleep(5 * time.Second)
}
