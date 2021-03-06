package engine_test

import (
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var cpu *engine.CPU

func TestMain(m *testing.M) {
	jb := engine.NewJamboy()
	cpu = jb.CPU

	os.Exit(m.Run())
}

func TestRegisters_SetMultiRegister(t *testing.T) {
	registers := GetTestRegisters()

	values := []uint{
		0x0000,
		0xFFFF,
		0xFF00,
		0x00FF,
		0x0101,
	}

	for _, r := range registers {
		for _, j := range values {
			cpu.WriteRegister(r.multiRegister, j)

			jh := uint8((j & 0xFF00) >> 8)
			jl := uint8(j & 0xFF)

			if r.multiRegister == engine.AF {
				jl = uint8(jl & 0xF0)
			}

			assert.Equal(t, jh, *r.a)
			assert.Equal(t, jl, *r.b)
		}
	}
}

func TestRegisters_GetMultiRegister(t *testing.T) {
	registers := GetTestRegisters()

	values := []uint8{
		0x00, 0x00,
		0xFF, 0xFF,
		0xFF, 0x00,
		0x00, 0xFF,
		0x01, 0x01,
	}

	for _, r := range registers {
		for j := 0; j < len(values); j += 2 {
			jh := values[j]
			jl := values[j+1]

			if r.multiRegister == engine.AF {
				jl = uint8(jl & 0xF0)
			}

			*r.a = jl
			*r.b = jh

			expected := (uint(jh) << 8) | uint(jl)

			value := cpu.ReadRegister(r.multiRegister)

			assert.Equal(t, value, expected)
		}
	}
}

func GetTestRegisters() []struct {
	multiRegister engine.RegisterID
	a             *uint8
	b             *uint8
} {
	registers := []struct {
		multiRegister engine.RegisterID
		a             *uint8
		b             *uint8
	}{
		{multiRegister: engine.AF, a: &cpu.Registers[engine.A], b: &cpu.Registers[engine.F]},
		{multiRegister: engine.BC, a: &cpu.Registers[engine.B], b: &cpu.Registers[engine.C]},
		{multiRegister: engine.DE, a: &cpu.Registers[engine.D], b: &cpu.Registers[engine.E]},
		{multiRegister: engine.HL, a: &cpu.Registers[engine.H], b: &cpu.Registers[engine.L]},
	}

	return registers
}
