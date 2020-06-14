package engine

import (
	"encoding/binary"
	"fmt"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLDd16(t *testing.T) {
	tests := []struct {
		r  engine.Register
		op engine.OpCode
	}{
		{
			r:  engine.BC,
			op: 0x01,
		},
		{
			r:  engine.DE,
			op: 0x11,
		},
		{
			r:  engine.HL,
			op: 0x21,
		},
		{
			r:  engine.SP,
			op: 0x31,
		},
	}

	for _, test := range tests {
		cart, jb := newTestJamboy()

		cart.Data = []byte{0xFF, 0x01}

		finished, err := engine.LDd16(jb, test.op)

		assert.True(t, finished)
		assert.Nil(t, err)

		expected := binary.BigEndian.Uint16(cart.Data)

		assert.Equal(t, expected, uint16(jb.CPU.ReadRegister(test.r)), "register value does not match expected")
	}
}

func TestLD(t *testing.T) {
	baseOp := uint8(0x40)

	tests := []struct {
		r  engine.Register
		op uint8
	}{
		{
			r:  engine.B,
			op: 0x00,
		},
		{
			r:  engine.C,
			op: 0x01,
		},
		{
			r:  engine.D,
			op: 0x02,
		},
		{
			r:  engine.E,
			op: 0x03,
		},
		{
			r:  engine.H,
			op: 0x04,
		},
		{
			r:  engine.L,
			op: 0x05,
		},
		{
			r:  engine.A,
			op: 0x07,
		},
	}

	var ldOrderedRegisters = []engine.Register{
		engine.B, engine.C, engine.D, engine.E, engine.H, engine.L, engine.Register(255), engine.A,
	}

	for loop := 0; loop < 8; loop++ {
		dstRegister := ldOrderedRegisters[loop]

		if dstRegister == 255 {
			for _, test := range tests {
				_, jb := newTestJamboy()

				jb.CPU.WriteRegister(engine.HL, 0xCCAA)
				jb.CPU.Registers[test.r] = 0xCC

				finished, err := engine.LD(jb, engine.OpCode(baseOp+test.op))

				assert.True(t, finished)
				assert.Nil(t, err)

				assert.Equal(
					t, uint8(0xCC), jb.Memory.RAM[jb.CPU.ReadRegister(engine.HL)],
					fmt.Sprintf("RAM does not match expected - src: %s dst: (HL)", test.r))
			}
		} else {
			for _, test := range tests {
				_, jb := newTestJamboy()

				jb.CPU.Registers[test.r] = 0xC0

				finished, err := engine.LD(jb, engine.OpCode(baseOp+test.op))

				assert.True(t, finished)
				assert.Nil(t, err)

				assert.Equal(
					t, uint(0xC0), jb.CPU.ReadRegister(dstRegister),
					fmt.Sprintf("register value does not match expected - src: %s dst: %s", test.r, dstRegister))
			}

			_, jb := newTestJamboy()

			jb.CPU.WriteRegister(engine.HL, 0x01)
			jb.Memory.Write(0x01, 0xC0)

			finished, err := engine.LD(jb, engine.OpCode(baseOp+0x06))

			assert.True(t, finished)
			assert.Nil(t, err)

			assert.Equal(
				t, uint(0xC0), jb.CPU.ReadRegister(dstRegister),
				fmt.Sprintf("register value does not match expected - src: %s dst: %s", engine.HL, dstRegister))

		}

		baseOp += 8
	}
}

func newTestJamboy() (*engine.Cart, *engine.Jamboy) {
	cart := &engine.Cart{
		Data: make([]byte, 0x100),
	}

	jb := engine.NewJamboy()

	jb.InsertCartridge(cart)

	jb.CPU.PC = 0x00
	return cart, jb
}
