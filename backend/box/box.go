package box

import (
	"github.com/carlmango11/schmarlbox/backend/box/bus"
	"github.com/carlmango11/schmarlbox/backend/box/cpu"
	"github.com/carlmango11/schmarlbox/backend/box/devices"
	"github.com/carlmango11/schmarlbox/backend/box/memory"
	"log"
)

const AddrDisplay = 0x2000
const AddrROM = 0x8000

type NES struct {
	tick int
	cpu  *cpu.CPU
}

func New(romData []byte) *NES {
	b := bus.New()

	rom := memory.New(AddrROM, romData)

	endROM := uint16(len(romData)) - 1
	endROM += AddrROM

	b.Connect(AddrDisplay, AddrDisplay, &devices.Display{})
	b.Connect(AddrROM, endROM, rom)

	return &NES{
		cpu: cpu.New(b),
	}
}

func (n *NES) Run() {
	var i int64
	//for range time.Tick(time.Second / cpu.ClockSpeedHz) {
	for {
		n.Tick()

		//log.Println("small tick")
		i++
		if i%(1660000) == 0 {
			log.Println("tick", i)
			i = 0
			//time.Sleep(10 * time.Millisecond)
		}
	}
}

func (n *NES) Tick() {
	n.cpu.Tick()
}

func (n *NES) Display() [25][80]byte {
	return [25][80]byte{}
}
