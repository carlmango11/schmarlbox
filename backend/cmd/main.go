package main

import (
	"github.com/carlmango11/schmarlbox/backend/box"
	"github.com/carlmango11/schmarlbox/backend/box/log"
	"golang.org/x/term"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Debug = true

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

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
}
