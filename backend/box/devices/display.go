package devices

import (
	"fmt"
)

type character byte

type Display struct {
	d [25][80]character
}

func (d *Display) Read(addr uint16) byte {
	return 0
}

func (d *Display) Write(addr uint16, val byte) {
	ch := character(val)

	d.handle(ch)
}

func (d *Display) handle(ch character) {
	fmt.Printf("%v", string(ch))
}
