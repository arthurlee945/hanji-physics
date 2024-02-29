package utility

import (
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/fzipp/canvas"
)

func StartCanvas(w, h int, title string, runFn func(*canvas.Context)) {
	http := flag.String("http", ":8080", "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, runFn, &canvas.Options{
		Title:          title,
		Width:          w,
		Height:         h,
		PageBackground: color.RGBA{R: 0xFA, G: 0xF9, B: 0xF6, A: 0xFF},
		EnabledEvents: []canvas.Event{
			canvas.MouseMoveEvent{},
		},
	})
	if err != nil {
		log.Fatalf("Failed on starting canvas server: %v", err)
	}
}

func httpLink(addr string) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	return "http://" + addr
}
