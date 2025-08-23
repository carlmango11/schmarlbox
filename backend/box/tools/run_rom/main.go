package main

import (
	"io"
	"os"

	"github.com/carlmango11/schmarlbox/backend/box"
	"github.com/carlmango11/schmarlbox/backend/box/log"
	"github.com/carlmango11/schmarlbox/backend/box/rom"
)

func main() {
	log.Debug = false

	f, err := os.Open("/Users/carl/IdeaProjects/Nes/backend/wasm/roms/donkey.nes")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	n := box.New(rom.FromBytes(bytes))
	n.Run()
}
