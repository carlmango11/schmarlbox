package main

import (
	"fmt"
	"github.com/carlmango11/schmarlbox/backend/box"
	"io"
	"os"
)

func main() {
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
