package box

import (
	"github.com/carlmango11/schmarlbox/backend/box/bus"
	"github.com/carlmango11/schmarlbox/backend/box/cpu"
	"github.com/carlmango11/schmarlbox/backend/box/devices"
	"github.com/carlmango11/schmarlbox/backend/box/keyboard"
	"github.com/carlmango11/schmarlbox/backend/box/memory"
	"time"
)

const (
	AddrZP       = 0x0000
	AddrStack    = 0x0100
	AddrRAM      = 0x0200
	AddrDisplay  = 0x2200
	AddrKeyboard = 0x3000
	AddrROM      = 0x8000
)

const (
	SizeZP    = 0x0100
	SizeStack = 0x0100
	SizeRAM   = 0x2000
)

type Box struct {
	tick    int
	cpu     *cpu.CPU
	display *devices.Display
}

func New(romData []byte) *Box {
	b := bus.New()

	rom := memory.New(AddrROM, len(romData), romData)
	endROM := uint16(len(romData)) - 1
	endROM += AddrROM

	display := devices.NewDisplay()

	stack := memory.New(AddrStack, SizeStack, nil)
	ram := memory.New(AddrRAM, SizeRAM, nil)

	b.Connect(AddrZP, SizeZP-1, memory.New(AddrZP, SizeZP, nil))
	b.Connect(AddrStack, AddrStack+SizeStack-1, stack)
	b.Connect(AddrRAM, AddrRAM+SizeRAM-1, ram)
	b.Connect(AddrDisplay, AddrDisplay, display)
	b.Connect(AddrKeyboard, AddrKeyboard, keyboard.New())
	b.Connect(AddrROM, endROM, rom)

	return &Box{
		cpu:     cpu.New(b),
		display: display,
	}
}

func (b *Box) Run() {
	for range time.Tick(time.Second / cpu.ClockSpeedHz) {
		//for {
		b.Tick()
	}
}

func (b *Box) Tick() {
	b.cpu.Tick()
}

func (b *Box) Display() [25][80]string {
	return b.display.State()
}
