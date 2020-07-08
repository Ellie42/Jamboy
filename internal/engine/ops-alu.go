package engine

import "fmt"

func SUB(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xD6 {
		jb.CPU.Subtract(A, jb.Read8Bit(), false)
		return
	}

	if opcode&0xF0 != 0x90 {
		panic(fmt.Sprintf("unhandled SUB %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.SubtractR8(A, register, false)

	return
}

func getOrderedRegister(opcode OpCode) RegisterID {
	return fullOrderedRegisters[(opcode&0x0F) % 8]
}

func SBC(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xDE {
		jb.CPU.Subtract(A, jb.Read8Bit(), true)
		return
	}

	if opcode&0xF0 != 0x90 {
		panic(fmt.Sprintf("unhandled SBC %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.SubtractR8(A, register, true)

	return nil
}

func ADC(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xCE {
		jb.CPU.Add(A, jb.Read8Bit(), true)
		return
	}

	if opcode&0xF0 != 0x80 {
		panic(fmt.Sprintf("unhandled ADC %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AddR8(A, register, true)

	return nil
}

func ADD(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xC6 {
		jb.CPU.Add(A, jb.Read8Bit(), false)
		return
	}

	if opcode&0xF0 != 0x80 {
		panic(fmt.Sprintf("unhandled ADD %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AddR8(A, register, false)

	return nil
}

func AND(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xE6 {
		jb.CPU.And(jb.Read8Bit())
		return
	}

	if opcode&0xF0 != 0xA0 {
		panic(fmt.Sprintf("unhandled AND %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AndR8(register)
	return
}

func XOR(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xEE {
		jb.CPU.Xor(jb.Read8Bit())
		return
	}

	if opcode&0xF0 != 0xA0 {
		panic(fmt.Sprintf("unhandled XOR %x", opcode))
	}

	jb.CPU.XorR8(getOrderedRegister(opcode))

	return nil
}

func OR(jb *Jamboy, opcode OpCode) (err error) {
	if opcode == 0xF6 {
		jb.CPU.Or(jb.Read8Bit())
		return
	}

	if opcode&0xF0 != 0xB0 {
		panic(fmt.Sprintf("unhandled OR %x", opcode))
	}

	jb.CPU.OrR8(getOrderedRegister(opcode))

	return
}

var addHLOrderedRegister = []RegisterID{
	BC, DE, HL, SP,
}

func ADDToHL(jb *Jamboy, opcode OpCode) (err error) {
	flags := Flag(0x00)

	hlVal := jb.CPU.ReadRegisterInstant(HL)
	other := jb.CPU.ReadRegister(addHLOrderedRegister[opcode&0xF0>>4])

	finalValue := hlVal + other

	if hlVal < 0x1000 && finalValue >= 0x1000 {
		flags |= HalfCarryFlag
	}

	if finalValue >= 0x10000 {
		flags |= CarryFlag
	}

	jb.CPU.WriteRegister(HL, finalValue)
	jb.CPU.SetFlagsMasked(flags, ZeroFlag|CarryFlag|HalfCarryFlag)

	return nil
}