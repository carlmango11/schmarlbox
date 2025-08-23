package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAddr(t *testing.T) {
	assert.Equal(t, uint16(0x2215), toAddr(0x22, 0x15))
}

func TestSTA_ZP(t *testing.T) {
	bus := &mockBus{}
	bus.Write(VectorReset, 0x00)
	bus.Write(VectorReset+1, 0x80)
	bus.Write(0x8000, 0x92)
	bus.Write(0x8001, 0x10)
	bus.Write(0x10, 0x77)
	bus.Write(0x11, 0x20)

	c := New(bus)
	c.a = 0x51

	c.Tick()

	assert.Equal(t, uint8(0x51), bus.Read(0x2077))
}

type mockBus struct {
	memory map[uint16]byte
}

func (b *mockBus) Read(a uint16) byte {
	return b.memory[a]
}

func (b *mockBus) Write(a uint16, v byte) {
	if b.memory == nil {
		b.memory = map[uint16]byte{}
	}

	b.memory[a] = v
}
