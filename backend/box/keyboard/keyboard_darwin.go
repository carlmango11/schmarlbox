package keyboard

import (
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
	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		key := buf[0]
		k.lastKey = key
	}
}

func (k *Keyboard) Read(addr uint16) byte {
	key := k.lastKey
	k.lastKey = 0

	return key
}

func (k *Keyboard) Write(addr uint16, val byte) {}
