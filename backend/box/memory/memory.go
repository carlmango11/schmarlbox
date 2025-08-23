package memory

type Memory struct {
	startAddr uint16
	data      []byte
}

func New(startAddr uint16, init []byte) *Memory {
	return &Memory{
		startAddr: startAddr,
		data:      init,
	}
}

func (m *Memory) Read(addr uint16) byte {
	return m.data[addr-m.startAddr]
}

func (m *Memory) Write(addr uint16, v byte) {

}
