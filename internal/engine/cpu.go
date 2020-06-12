package engine

// REMEMBER LITTLE ENDIAN
type CPU struct {
	PC uint16
	SP uint16

	Registers
}

//go:generate stringer -type=MultiRegister -linecomment
type MultiRegister int

const (
	AF MultiRegister = iota // AF
	BC                      // BC
	DE                      // DE
	HL                      // HL
)

type Registers struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8

	// Holds CPU flags
	// Lower nibble always 0
	F uint8
	H uint8
	L uint8
}

func (r *Registers) SetMultiRegister(multi MultiRegister, value uint16) {
	switch multi {
	case AF:
		r.A = uint8((value & 0xFF00) >> 8)
		r.F = uint8(value & 0x00FF)
	case BC:
		r.B = uint8((value & 0xFF00) >> 8)
		r.C = uint8(value & 0x00FF)
	case DE:
		r.D = uint8((value & 0xFF00) >> 8)
		r.E = uint8(value & 0x00FF)
	case HL:
		r.H = uint8((value & 0xFF00) >> 8)
		r.L = uint8(value & 0x00FF)
	}
}

func (r Registers) GetMultiRegister(multi MultiRegister) (value uint16) {
	switch multi {
	case AF:
		value = (uint16(r.A) << 8) | uint16(r.F)
	case BC:
		value = (uint16(r.B) << 8) | uint16(r.C)
	case DE:
		value = (uint16(r.D) << 8) | uint16(r.E)
	case HL:
		value = (uint16(r.H) << 8) | uint16(r.L)
	}

	return
}

func (CPU) ExecuteOp() {

}
