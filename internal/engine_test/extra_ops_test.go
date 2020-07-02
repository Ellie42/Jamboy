package engine

import (
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBITOperations(t *testing.T) {
	jb := engine.NewJamboy()

	for bit := 0; bit < 8; bit++ {
		jb.CPU.WriteRegister(engine.B, 0x00)

		err := engine.BIT(jb, engine.OpCode(0x40+bit*8))

		if err != nil {
			t.Fail()
		}

		assert.Equal(t, uint8(engine.ZeroFlag), uint8(jb.CPU.GetFlags()&engine.ZeroFlag))

		jb.CPU.WriteRegister(engine.B, 1<<bit)

		err = engine.BIT(jb, engine.OpCode(0x40+bit*8))

		if err != nil {
			t.Fail()
		}

		assert.Equal(t, uint8(0), uint8(jb.CPU.GetFlags()&engine.ZeroFlag))
	}
}
