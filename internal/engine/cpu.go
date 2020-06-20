package engine

import "fmt"

// REMEMBER LITTLE ENDIAN
type CPU struct {
	PC uint16
	SP uint16

	Registers []uint8

	memory           *MMU
	CurrentJumpTable *[]func(jb *Jamboy, opcode OpCode) (err error)
	CurrentOP        OpCode
	Jamboy           *Jamboy
}

//go:generate stringer -type=Register -linecomment
type RegisterID int

const (
	A  RegisterID = iota // A
	B                    // B
	C                    // C
	D                    // D
	E                    // E
	F                    // F
	H                    // H
	L                    // L
	AF                   // AF
	BC                   // BC
	DE                   // DE
	HL                   // HL
	SP                   // SP
	PC                   // PC
)

type Register byte

type Registers struct {
	A Register
	B Register
	C Register
	D Register
	E Register

	// Holds CPU flags
	// Lower nibble always 0
	F Register
	H Register
	L Register
}

func (r Register) String() string {
	return fmt.Sprintf("0x%02x", r)
}

type Flag byte

const (
	ZeroFlag Flag = 1 << (7 - iota)
	SubFlag
	HalfCarryFlag
	CarryFlag
)

func (c *CPU) WriteRegister(register RegisterID, value uint) {
	c.Wait(1)
	c.WriteRegisterInstant(register, value)
}

func (c *CPU) WriteRegisterInstant(register RegisterID, value uint) {
	switch register {
	case AF:
		c.Registers[F] = byte(value & 0x00F0)
		c.Registers[A] = byte((value & 0xFF00) >> 8)
	case BC:
		c.Registers[C] = byte(value & 0x00FF)
		c.Registers[B] = byte((value & 0xFF00) >> 8)
	case DE:
		c.Registers[E] = byte(value & 0x00FF)
		c.Registers[D] = byte((value & 0xFF00) >> 8)
	case HL:
		c.Registers[L] = byte(value & 0x00FF)
		c.Registers[H] = byte((value & 0xFF00) >> 8)
	case A:
		c.Registers[register] = byte(value)
	case B:
		c.Registers[register] = byte(value)
	case C:
		c.Registers[register] = byte(value)
	case D:
		c.Registers[register] = byte(value)
	case E:
		c.Registers[register] = byte(value)
	case F:
		c.Registers[register] = byte(value & 0xF0)
	case H:
		c.Registers[register] = byte(value)
	case L:
		c.Registers[register] = byte(value)
	case SP:
		c.SP = uint16(value)
	case PC:
		c.PC = uint16(value)
	}
}

func (c *CPU) ReadRegister(register RegisterID) uint {
	c.Wait(1)

	return c.ReadRegisterInstant(register)
}

func (c *CPU) ReadRegisterInstant(register RegisterID) (value uint) {
	switch register {
	case AF:
		value = uint(c.Registers[F])
		value |= (uint(c.Registers[A]) << 8)
	case BC:
		value = uint(c.Registers[C])
		value |= (uint(c.Registers[B]) << 8)
	case DE:
		value = uint(c.Registers[E])
		value |= (uint(c.Registers[D]) << 8)
	case HL:
		value = uint(c.Registers[L])
		value |= (uint(c.Registers[H]) << 8)
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

func (c *CPU) AddFlags(flags Flag) {
	c.WriteRegisterInstant(F, c.ReadRegisterInstant(F)|uint(flags))
}

func (c *CPU) SetFlags(flags Flag) {
	c.WriteRegisterInstant(F, uint(flags))
}

func (c *CPU) GetFlags() Flag {
	return Flag(c.ReadRegister(F))
}

func NewCPU(jb *Jamboy) *CPU {
	return &CPU{
		Jamboy:    jb,
		Registers: make([]uint8, 8),
	}
}