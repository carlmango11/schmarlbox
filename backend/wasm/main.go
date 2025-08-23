package main

import (
	_ "embed"
	"syscall/js"

	"github.com/carlmango11/schmarlbox/backend/box"
	"github.com/carlmango11/schmarlbox/backend/box/log"
)

//go:embed bios.out
var rom []byte

func main() {
	createBindings()

	waitC := make(chan bool)
	<-waitC
}

func createBindings() {
	var b *box.Box

	log.Debug = false

	b = box.New(rom)
	log.Printf("DONE %v", len(rom))
	go b.Run()

	getDisplayFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if b == nil {
			return 1
		}

		display := b.Display()

		height := len(display)
		width := len(display[0])

		result := make([]any, height*width)

		// doesn't support normal 2D typed arrays, only []any
		for y := range display {
			for x := range display[y] {
				result[x+(y*width)] = display[y][x]
			}
		}

		return result
	})

	js.Global().Set("getDisplay", getDisplayFunc)
}
