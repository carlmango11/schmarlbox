package keyboard

import (
	"github.com/eiannone/keyboard"
	"log"
	"os"
)

type Keyboard struct {
	lastKey byte
}

func New() *Keyboard {
	k := &Keyboard{}
	go k.listen()

	return k
}

func (k *Keyboard) listen() {
	if os.Getenv("IDE") == "true" {
		return
	}

	if err := keyboard.Open(); err != nil {
		log.Panicf("keyboard not started: %v", err)
	}
	defer keyboard.Close()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Panicf("keyboard err: %v", err)
		}

		k.lastKey = byte(char)
	}
}

func (k *Keyboard) Read(addr uint16) byte {
	key := k.lastKey
	k.lastKey = 0

	return key
}

func (k *Keyboard) Write(addr uint16, val byte) {}
