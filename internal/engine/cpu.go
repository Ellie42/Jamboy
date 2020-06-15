package engine

// REMEMBER LITTLE ENDIAN
type CPU struct {
	PC uint16
	SP uint16

	Registers []uint8

	memory           *Memory
	CurrentJumpTable *[]func(jb *Jamboy, opcode OpCode) (finished bool, err error)
	CurrentOP        OpCode
}

//go:generate stringer -type=Register -linecomment
type Register int

const (
	A  Register = iota // A
	B                  // B
	C                  // C
	D                  // D
	E                  // E
	F                  // F
	H                  // H
	L                  // L
	AF                 // AF
	BC                 // BC
	DE                 // DE
	HL                 // HL
	SP                 // SP
	PC                 // PC
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

type Flag byte

const (
	ZeroFlag      Flag = 1 << (7 - iota)
	SubFlag
	HalfCarryFlag
	CarryFlag
)

func (c *CPU) WriteRegister(register Register, value uint) {
	c.Wait(1)
	c.WriteRegisterInstant(register, value)
}

func (c *CPU) WriteRegisterInstant(register Register, value uint) {
	switch register {
	case AF:
		c.Registers[A] = uint8(value & 0x00FF)
		c.Registers[F] = uint8((value & 0xFF00) >> 8)
	case BC:
		c.Registers[B] = uint8(value & 0x00FF)
		c.Registers[C] = uint8((value & 0xFF00) >> 8)
	case DE:
		c.Registers[D] = uint8(value & 0x00FF)
		c.Registers[E] = uint8((value & 0xFF00) >> 8)
	case HL:
		c.Registers[H] = uint8(value & 0x00FF)
		c.Registers[L] = uint8((value & 0xFF00) >> 8)
	case A:
		fallthrough
	case B:
		fallthrough
	case C:
		fallthrough
	case D:
		fallthrough
	case E:
		fallthrough
	case F:
		fallthrough
	case H:
		fallthrough
	case L:
		c.Registers[register] = uint8(value)
	case SP:
		c.SP = uint16(value)
	case PC:
		c.PC = uint16(value)
	}
}

func (c *CPU) ReadRegister(register Register) uint {
	c.Wait(1)

	return c.ReadRegisterInstant(register)
}

func (c *CPU) ReadRegisterInstant(register Register) (value uint) {
	switch register {
	case AF:
		value = uint(c.Registers[A])
		value |= (uint(c.Registers[F]) << 8)
	case BC:
		value = uint(c.Registers[B])
		value |= (uint(c.Registers[C]) << 8)
	case DE:
		value = uint(c.Registers[D])
		value |= (uint(c.Registers[E]) << 8)
	case HL:
		value = uint(c.Registers[H])
		value |= (uint(c.Registers[L]) << 8)
	case A:
		fallthrough
	case B:
		fallthrough
	case C:
		fallthrough
	case D:
		fallthrough
	case E:
		fallthrough
	case F:
		fallthrough
	case H:
		fallthrough
	case L:
		value = uint(c.Registers[register])
	case SP:
		value = uint(c.SP)
	case PC:
		value = uint(c.PC)
	}

	return
}

func (CPU) ExecuteOp() {

}

func (c *CPU) Reset() {
	c.CurrentJumpTable = &BaseOpJumpTable

	c.WriteRegister(AF, 0x01B0)
	c.WriteRegister(BC, 0x0013)
	c.WriteRegister(DE, 0x00D8)
	c.WriteRegister(HL, 0x014D)
	c.SP = 0xFFFE
	c.PC = 0x0100
}

func (c CPU) Wait(i int) {

}

func (c *CPU) IncrementHL() {
	c.WriteRegisterInstant(HL, c.ReadRegisterInstant(HL)+1)
}

func (c *CPU) DecrementHL() {
	c.WriteRegisterInstant(HL, c.ReadRegisterInstant(HL)-1)
}

func (c *CPU) AddFlags(flags Flag){
	c.WriteRegisterInstant(AF, c.ReadRegisterInstant(AF) | uint(flags))
}

func (c *CPU) SetFlags(flags Flag){
	c.WriteRegisterInstant(AF, uint(flags))
}

func NewCPU(memory *Memory) *CPU {
	return &CPU{
		memory:    memory,
		Registers: make([]uint8, 8),
	}
}