package main

import (
	"fmt"
	"github.com/carlmango11/schmarlbox/backend/box"
	"github.com/carlmango11/schmarlbox/backend/box/log"
	"io"
	"os"
)

func main() {
	if os.Getenv("IDE") == "true" {
		log.Debug = true
	}

	f, err := os.Open("/Users/carl/IdeaProjects/schmarlbox/build/bios.out")
	if err != nil {
		panic(err)
	}

	romData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	box := box.New(romData)

	box.Run()

	fmt.Println("END")
}
