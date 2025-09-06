package cpu

func (c *CPU) initBranch() {
	instrs := map[byte]Instr{
		// Branch
		0x10: {name: "BPL", addrMode: Relative, condition: c.bpl},
		0x30: {name: "BMI", addrMode: Relative, condition: c.bmi},
		0x50: {name: "BVC", addrMode: Relative, condition: c.bvc},
		0x70: {name: "BVS", addrMode: Relative, condition: c.bvs},
		0x90: {name: "BCC", addrMode: Relative, condition: c.bcc},
		0xB0: {name: "BCS", addrMode: Relative, condition: c.bcs},
		0xD0: {name: "BNE", addrMode: Relative, condition: c.bne},
		0xF0: {name: "BEQ", addrMode: Relative, condition: c.beq},
	}

	for code, instr := range instrs {
		c.opCodes[code] = instr
	}
}

func (c *CPU) bpl() bool {
	return !c.controlFlagSet(FlagN)
}

func (c *CPU) bmi() bool {
	return c.controlFlagSet(FlagN)
}

func (c *CPU) bvc() bool {
	return !c.controlFlagSet(FlagV)
}

func (c *CPU) bvs() bool {
	return c.controlFlagSet(FlagV)
}

func (c *CPU) bcc() bool {
	return !c.controlFlagSet(FlagC)
}

func (c *CPU) bcs() bool {
	return c.controlFlagSet(FlagC)
}

func (c *CPU) bne() bool {
	return !c.controlFlagSet(FlagZ)
}

func (c *CPU) beq() bool {
	return c.controlFlagSet(FlagZ)
}
