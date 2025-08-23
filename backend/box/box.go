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
	AddrKeyboard = 0x3000
	AddrStack    = 0x0100
	AddrDisplay  = 0x2000
	AddrROM      = 0x8000
	AddrZP       = 0x0000
)

const (
	SizeStack = 0x0100
	SizeZP    = 0x0100
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

	b.Connect(AddrZP, SizeZP-1, memory.New(AddrZP, SizeZP, nil))
	b.Connect(AddrDisplay, AddrDisplay, display)
	b.Connect(AddrROM, endROM, rom)
	b.Connect(AddrStack, AddrStack+SizeStack-1, stack)
	b.Connect(AddrKeyboard, AddrKeyboard, keyboard.New())

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
