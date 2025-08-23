package main

import (
	_ "embed"
	"time"

	"github.com/carlmango11/schmarlbox/backend/box"
	"github.com/carlmango11/schmarlbox/backend/box/log"
	"github.com/carlmango11/schmarlbox/backend/box/rom"
)

//go:embed donkey1.nes
var donkeyRom []byte

//go:embed color_test.nes
var colourRom []byte

func main() {
	log.Debug = false

	n := box.New(rom.FromBytes(donkeyRom))
	go n.Run()

	for range time.Tick(time.Second) {
		n.Display()
	}
}
