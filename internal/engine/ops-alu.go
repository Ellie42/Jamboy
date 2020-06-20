package engine

import "fmt"

func SUB(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0x90 {
		panic(fmt.Sprintf("unhandled SUB %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.SubtractR8(A, register, false)

	return nil
}

func getOrderedRegister(opcode OpCode) RegisterID {
	return fullOrderedRegisters[(opcode&0x0F) % 8]
}

func SBC(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0x90 {
		panic(fmt.Sprintf("unhandled SBC %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.SubtractR8(A, register, true)

	return nil
}

func ADC(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0x80 {
		panic(fmt.Sprintf("unhandled ADC %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AddR8(A, register, true)

	return nil
}

func ADD(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0x80 {
		panic(fmt.Sprintf("unhandled ADD %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AddR8(A, register, false)

	return nil
}

func AND(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0xA0 {
		panic(fmt.Sprintf("unhandled AND %x", opcode))
	}

	register := getOrderedRegister(opcode)

	jb.CPU.AndR8(register)

	panic(fmt.Sprintf("not implemented op AND -  %x", opcode))
}

func XOR(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0xA0 {
		panic(fmt.Sprintf("unhandled XOR %x", opcode))
	}

	jb.CPU.XorR8(getOrderedRegister(opcode))

	return nil
}

func OR(jb *Jamboy, opcode OpCode) (err error) {
	if opcode&0xF0 != 0xB0 {
		panic(fmt.Sprintf("unhandled OR %x", opcode))
	}

	jb.CPU.OrR8(getOrderedRegister(opcode))

	return nil
}
