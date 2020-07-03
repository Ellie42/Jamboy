package engine

import (
	"fmt"
)

// REMEMBER LITTLE ENDIAN
type CPU struct {
	PC uint16
	SP uint16

	Registers []uint8

	CurrentJumpTable *[]func(jb *Jamboy, opcode OpCode) (err error)
	CurrentOP        OpCode
	Jamboy           *Jamboy
	Cycles           uint
	IsBooting        bool
	HasBootROM       bool
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
	case F:
		c.Registers[register] = byte(value & 0xF0)
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
	case H:
		fallthrough
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

	if c.HasBootROM {
		c.IsBooting = true
		c.PC = 0x0000
	} else {
		c.IsBooting = false
		c.PC = 0x0100
	}
}

func (c CPU) Wait(i uint) {
	c.Cycles += i * 4
	c.Jamboy.GPU.Clocks += i
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

func (c *CPU) SetFlagsMasked(flags Flag, mask Flag) {
	c.WriteRegisterInstant(F, uint((c.GetFlags() & ^mask)|flags))
}

func (c *CPU) GetFlags() Flag {
	return Flag(c.ReadRegister(F))
}

func (c *CPU) IsFlagSet(flag Flag) bool {
	return c.GetFlags()&flag > 0
}

func (c *CPU) Call(pointer uint16) {
	returnAddr := c.PC

	c.PushUint16(returnAddr)

	c.Jump(pointer)
}

func (c *CPU) PushUint16(data uint16) {
	c.Push(byte((data & 0xFF00 >> 8)))
	c.Push(byte((data & 0x00FF)))
}

func (c *CPU) Push(data byte) {
	c.SP -= 1
	c.Jamboy.MMU.Write(Address(c.SP), data)
}

func (c *CPU) Pop() (data byte) {
	data = c.Jamboy.MMU.Read(Address(c.SP))
	c.SP += 1
	return
}

func (c *CPU) PopUint16() (data uint16) {
	data = uint16(c.Pop())
	data |= uint16(c.Pop()) << 8
	return
}

func (c *CPU) Jump(pointer uint16) {
	c.WriteRegister(PC, uint(pointer))
}

func (c *CPU) Return() {
	c.Jump(c.PopUint16())
}

func (c *CPU) AddR8(a RegisterID, b RegisterID, withCarry bool) {
	c.Add(a, c.getRegisterValue(b), withCarry)
}

func (c *CPU) Add(a RegisterID, b uint8, withCarry bool) {
	srcValue := c.ReadRegister(a)
	finalValue := uint(srcValue)

	if withCarry && c.IsFlagSet(CarryFlag) {
		finalValue += 1
	}

	c.SetFlags(Flag(c.Registers[F] & 0x10))

	if a == HL {
		hl := srcValue
		srcValue = uint(c.Jamboy.MMU.Read(Address(hl)))
		finalValue = srcValue + uint(b)
		c.Jamboy.MMU.Write(Address(hl), byte(finalValue))
	} else {
		finalValue = srcValue + uint(b)
		c.WriteRegister(a, finalValue)
	}

	if finalValue >= 0x100 {
		c.AddFlags(ZeroFlag)
	}

	if srcValue < 0x10 && finalValue >= 0x10 {
		c.AddFlags(HalfCarryFlag)
	}

	if finalValue > 0xFF {
		c.AddFlags(CarryFlag)
	}
}

func (c *CPU) SubtractR8(a RegisterID, b RegisterID, withCarry bool) {
	c.Subtract(a, c.getRegisterValue(b), withCarry)
}

func (c *CPU) getRegisterValue(register RegisterID) (value uint8) {
	if register == HL {
		value = c.Jamboy.MMU.Read(Address(c.ReadRegisterInstant(register)))
	} else {
		value = uint8(c.ReadRegister(register))
	}

	return
}

func (c *CPU) Subtract(a RegisterID, b uint8, withCarry bool) {
	srcValue := int(c.ReadRegister(a))
	finalValue := int(srcValue)

	if withCarry && c.IsFlagSet(CarryFlag) {
		finalValue -= 1
	}

	c.SetFlags(Flag(c.Registers[F]&0x10) | SubFlag)

	if a == HL {
		hl := srcValue
		srcValue = int(c.Jamboy.MMU.Read(Address(hl)))
		finalValue = srcValue - int(b)
		c.Jamboy.MMU.Write(Address(hl), byte(finalValue))
	} else {
		finalValue = srcValue - int(b)
		c.WriteRegister(a, uint(finalValue))
	}

	if finalValue == 0 {
		c.AddFlags(ZeroFlag)
	}

	if srcValue >= 0x10 && finalValue < 0x10 {
		c.AddFlags(HalfCarryFlag)
	}

	if finalValue < 0 {
		c.AddFlags(CarryFlag)
	}
}

func (c *CPU) AndR8(register RegisterID) {
	c.And(c.getRegisterValue(register))
}

func (c *CPU) And(value uint8) {
	c.SetFlags(0x00)

	finalValue := value & c.Registers[A]

	if finalValue == 0 {
		c.AddFlags(ZeroFlag)
	}

	c.Registers[A] = finalValue
}

func (c *CPU) OrR8(register RegisterID) {
	c.Or(c.getRegisterValue(register))
}

func (c *CPU) Or(value uint8) {
	c.SetFlags(0x00)

	finalValue := value | c.Registers[A]

	if finalValue == 0 {
		c.AddFlags(ZeroFlag)
	}

	c.Registers[A] = finalValue
}

func (c *CPU) XorR8(register RegisterID) {
	c.Or(c.getRegisterValue(register))
}

func (c *CPU) Xor(value uint8) {
	c.SetFlags(0x00)

	finalValue := value ^ c.Registers[A]

	if finalValue == 0 {
		c.AddFlags(ZeroFlag)
	}

	c.Registers[A] = finalValue
}

func (c *CPU) CompareR8(register RegisterID) {
	c.Compare(c.getRegisterValue(register))
}

func (c *CPU) Compare(value uint8) {
	c.SetFlags(0x00)

	if c.Registers[A] >= 0x10 && value > 0 {
		c.AddFlags(HalfCarryFlag)
	}

	if value > c.Registers[A] {
		c.AddFlags(CarryFlag)
	}

	if value == c.Registers[A] {
		c.AddFlags(ZeroFlag)
	}
}

func (c *CPU) GetFlag(flag Flag) uint8 {
	if c.GetFlags()&flag > 0 {
		return 1
	}

	return 0
}

func (c *CPU) LoadBootRom(data []byte) {
	c.HasBootROM = true
	copy(c.Jamboy.MMU.RAM[0:256], data)
}

func NewCPU(jb *Jamboy) *CPU {
	return &CPU{
		Jamboy:    jb,
		Registers: make([]uint8, 8),
	}
}
