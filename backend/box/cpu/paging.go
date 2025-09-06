package cpu

func (c *CPU) initPaging() {
	instrs := map[byte]Instr{
		0x02: {
			name:   "PGE",
			cycles: 1,
			flagChange: &flagChange{
				target: TargetPaging,
				flag:   FlagPaging,
				set:    true,
			},
		},
		0x03: {
			name:   "PGD",
			cycles: 1,
			flagChange: &flagChange{
				target: TargetPaging,
				flag:   FlagPaging,
				set:    false,
			},
		},
	}

	for code, instr := range instrs {
		c.opCodes[code] = instr
	}
}
