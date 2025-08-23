package bus

import (
	"fmt"
	"sort"

	"github.com/carlmango11/schmarlbox/backend/box/log"
)

type Component interface {
	Read(uint16) byte
	Write(uint16, byte)
}

type ComponentEntry struct {
	start     uint16
	end       uint16
	component Component
}

type Bus struct {
	components []*ComponentEntry
}

func New() *Bus {
	return &Bus{}
}

func (b *Bus) Connect(start, end uint16, c Component) {
	log.Printf("connected from 0x%x to 0x%x", start, end)

	b.components = append(b.components, &ComponentEntry{
		start:     start,
		end:       end,
		component: c,
	})

	sort.Slice(b.components, func(i, j int) bool {
		return b.components[i].start < b.components[j].start
	})
}

func (b *Bus) Read(addr uint16) byte {
	log.Debugf("bus: read 0x%x", addr)
	return b.getComponent(addr).Read(addr)
}

func (b *Bus) Write(addr uint16, v byte) {
	log.Debugf("bus: write at 0x%x: 0x%x", addr, v)
	b.getComponent(addr).Write(addr, v)
}

func (b *Bus) getComponent(addr uint16) Component {
	if addr < 0 {
		panic(fmt.Sprintf("bus: read from invalid address %x", addr))
	}

	for _, c := range b.components {
		if addr >= c.start && addr <= c.end {
			return c.component
		}
	}

	panic(fmt.Sprintf("bus: unhandled address %x", addr))
}
