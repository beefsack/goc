package main

import (
	"log"

	"github.com/beefsack/goc"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatalf("failed to init termbox: %s", err)
	}
	defer termbox.Close()
	goc.Screen(goc.PaddedEven(&goc.Bordered{
		Content: &goc.Text{
			Content: goc.StrToCells(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin sit amet dui odio. Aenean rutrum quam non augue gravida vehicula. Mauris justo ex, tincidunt a neque nec, sodales sodales odio. Fusce eget condimentum ex, id cursus enim. Donec eget eros eleifend, porta nulla et, volutpat nisi. Nunc ultrices non eros vel volutpat. Fusce molestie dolor elementum, faucibus massa ac, maximus ipsum. Sed rhoncus non libero ut iaculis. Nam eget accumsan est. Phasellus id nunc odio. Donec dignissim vitae odio nec porta. Morbi at congue ex. Quisque diam massa, bibendum id ante ut, bibendum mattis ipsum. Morbi pellentesque turpis leo, et condimentum nulla mattis aliquam. Nunc libero velit, fermentum dictum interdum a, consequat non ex.

Maecenas quis ipsum sapien. Nulla facilisi. Nulla feugiat mauris id maximus ultrices. Vestibulum nec sodales arcu. Sed vehicula lectus ac lorem posuere fringilla. Praesent ut ante quis justo euismod eleifend tristique at nulla. Proin facilisis augue tortor, eget porta justo posuere elementum. Suspendisse diam lectus, placerat id quam ac, iaculis porttitor magna. Cras egestas leo eros, nec mattis nisi rhoncus eu. Proin auctor ante dolor, at dignissim nulla vehicula sed. Donec sodales in ex sit amet blandit. Donec ultrices justo facilisis leo elementum tristique. Pellentesque condimentum lobortis sem eu egestas. Curabitur scelerisque, justo vitae euismod euismod, turpis tellus sollicitudin risus, quis suscipit purus neque vel felis. Quisque ac malesuada arcu. Donec a mi tristique, mattis turpis in, sagittis magna.

Vestibulum vitae ligula vitae neque pharetra cursus sit amet at magna. Donec gravida blandit metus sed porta. Aliquam ultrices nunc sed neque tempor fringilla. Suspendisse eget sem eget mauris interdum egestas. Vestibulum cursus tincidunt dolor, ac cursus ligula egestas sit amet. Sed malesuada eget mi non aliquet. Curabitur sed nunc in leo rhoncus faucibus a non ante. Ut congue massa ex, sit amet tempor magna rhoncus ut. Mauris sagittis ultricies orci, nec malesuada urna.

Duis et metus vitae sem lobortis tristique. Vestibulum ornare euismod ullamcorper. Sed venenatis quam lacus, non convallis augue volutpat vel. Praesent ut neque ut sapien pellentesque aliquam ut eu tellus. Etiam nec turpis ex. Etiam luctus porttitor sollicitudin. Vivamus ac ex nec quam semper condimentum nec id orci. Sed viverra posuere lorem. Nunc varius ante ac lectus eleifend, non placerat arcu malesuada. Curabitur vitae sapien urna. Integer id malesuada lorem. Quisque tellus ex, maximus non finibus et, elementum a dui. Nullam vitae lectus tortor.

Fusce volutpat metus sit amet aliquet venenatis. Nam quis ante maximus, consectetur lectus ut, efficitur justo. Ut tempus, arcu vitae laoreet viverra, ex felis aliquet magna, cursus auctor nunc ligula vitae odio. In hendrerit dui enim, vel consequat elit mattis quis. Suspendisse volutpat nibh nulla, quis ullamcorper diam hendrerit in. In sit amet venenatis massa. Suspendisse potenti. Curabitur rutrum finibus fringilla. Praesent molestie, libero ut finibus tincidunt, neque mauris tristique ante, quis posuere elit nisi et neque. Curabitur fringilla maximus volutpat.`, termbox.ColorRed, termbox.ColorDefault),
		},
		Label:    "Check out the border!",
		BorderFg: termbox.ColorGreen | termbox.AttrBold,
		LabelFg:  termbox.ColorBlue | termbox.AttrUnderline | termbox.AttrBold,
	}, 1))
	if err := termbox.Flush(); err != nil {
		log.Fatalf("failed to flush, %s", err)
	}
	termbox.PollEvent()
}
