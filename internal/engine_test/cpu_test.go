package engine_test

import (
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"os"
	"testing"
)

var cpu *engine.CPU

func TestMain(m *testing.M) {
	cpu = &engine.CPU{}

	os.Exit(m.Run())
}

func TestRegisters_SetMultiRegister(t *testing.T) {
	registers := GetTestRegisters()

	values := []uint16{
		0x0000,
		0xFFFF,
		0xFF00,
		0x00FF,
		0x0101,
	}

	for _, r := range registers {
		for _, j := range values {
			cpu.SetMultiRegister(r.multiRegister, j)

			jl := uint8((j & 0xFF00) >> 8)
			jh := uint8(j & 0xFF)

			if *r.a != jl || *r.b != jh {
				t.Errorf(
					"Invalid register values for %s: %d, %d\nExpected: %d, %d",
					r.multiRegister.String(), *r.a, *r.b, jl, jh,
				)
			}
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
			jl := values[j]
			jh := values[j+1]

			*r.a = jl
			*r.b = jh

			expected := (uint16(jl) << 8) | uint16(jh)

			value := cpu.GetMultiRegister(r.multiRegister)

			if value != expected {
				t.Errorf(
					"Returned value not correct for %s: %x\nExpected: %x",
					r.multiRegister.String(), value, expected,
				)
			}
		}
	}
}

func GetTestRegisters() []struct {
	multiRegister engine.MultiRegister
	a             *uint8
	b             *uint8
} {
	registers := []struct {
		multiRegister engine.MultiRegister
		a             *uint8
		b             *uint8
	}{
		{multiRegister: engine.AF, a: &cpu.A, b: &cpu.F},
		{multiRegister: engine.BC, a: &cpu.B, b: &cpu.C},
		{multiRegister: engine.DE, a: &cpu.D, b: &cpu.E},
		{multiRegister: engine.HL, a: &cpu.H, b: &cpu.L},
	}

	return registers
}
