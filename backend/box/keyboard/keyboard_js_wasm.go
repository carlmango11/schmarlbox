package keyboard

type Keyboard struct {
	lastKey byte
}

func New() *Keyboard {
	k := &Keyboard{}

	return k
}

func (k *Keyboard) Read(addr uint16) byte {
	key := k.lastKey
	k.lastKey = 0

	return key
}

func (k *Keyboard) Write(addr uint16, val byte) {}
