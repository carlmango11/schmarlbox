package pins

import (
	"fmt"
	"github.com/carlmango11/schmarlbox/backend/box/vectors"
)

type Pins struct {
	interruptLo byte
	interruptHi byte
}

func (p *Pins) Read(addr uint16) byte {
	switch {
	case addr == vectors.Reset:
		return 0
	case addr == vectors.Reset+1:
		return 0
	case addr == vectors.IRQ:
		return p.interruptLo
	case addr == vectors.IRQ+1:
		return p.interruptHi
	default:
		panic(fmt.Sprintf("bus: read from unhandled address %x", addr))
	}
}

func (p *Pins) Write(addr uint16, v byte) {
	switch {
	case addr == vectors.IRQ:
		p.interruptLo = v
	case addr == vectors.IRQ+1:
		p.interruptHi = v
	default:
		panic(fmt.Sprintf("bus: read from unhandled address %x", addr))
	}
}
