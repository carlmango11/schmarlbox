package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAddr(t *testing.T) {
	assert.Equal(t, uint16(0x2215), toAddr(0x22, 0x15))
}
