package memory

type Memory struct {
	startAddr uint16
	data      []byte
}

func New(startAddr uint16, size int, init []byte) *Memory {
	data := make([]byte, size, size)

	for i, b := range init {
		data[i] = b
	}

	return &Memory{
		startAddr: startAddr,
		data:      data,
	}
}

func (m *Memory) Read(addr uint16) byte {
	return m.data[addr-m.startAddr]
}

func (m *Memory) Write(addr uint16, v byte) {
	m.data[addr-m.startAddr] = v
}
