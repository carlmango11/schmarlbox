package cpu

import (
	"os"
)

func (c *CPU) initCtrl() {
	instrs := map[byte]Instr{
		0x00: {
			name:           "BRK",
			cycles:         7,
			impliedHandler: c.brk,
			addrMode:       Implied,
		},
		0x4C: {
			name:        "JMP",
			cycles:      3,
			addrHandler: c.jmp,
			addrMode:    AbsoluteAddr,
		},
		0x6C: {
			name:        "JMP",
			cycles:      5,
			addrHandler: c.jmp,
			addrMode:    Indirect,
		},
		0x20: {
			name:        "JSR",
			cycles:      5,
			addrHandler: c.jsr,
			addrMode:    AbsoluteAddr,
		},
		0x60: {
			name:           "RTS",
			cycles:         6,
			impliedHandler: c.rts,
			addrMode:       Implied,
		},
		0x40: {
			name:           "RTI",
			cycles:         6,
			impliedHandler: c.rti,
			addrMode:       Implied,
		},
	}

	for code, instr := range instrs {
		c.opCodes[code] = instr
	}
}

func (c *CPU) brk() {
	//c.PrintState()
	//return
	//log.Printf("breakpoint")
	//return
	if c.a == 0 {
		os.Exit(0)
	}
	panic(c.a)
	os.Exit(int(c.a))
	c.pushAddr(c.pc + 1)
	c.pushFlagsToStack()

	lo := c.bus.Read(VectorIRQ)
	hi := c.bus.Read(VectorIRQ + 1)

	c.pc = toAddr(hi, lo)
}

func (c *CPU) rti() {
	c.p = c.popStack()
	lo := c.popStack()
	hi := c.popStack()

	c.pc = toAddr(hi, lo)
}

func (c *CPU) jmp(addr uint16) {
	c.pc = addr
	// TODO: do I need to implement the weird behaviour around end of page?
}

func (c *CPU) jsr(addr uint16) {
	c.pushAddr(c.pc - 1)
	c.pc = addr
}

func (c *CPU) rts() {
	lo := c.popStack()
	hi := c.popStack()

	c.pc = toAddr(hi, lo) + 1
}
