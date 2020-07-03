package engine

import "fmt"

func LDAToRAMPointer(jb *Jamboy, opcode OpCode) (err error) {
	pointer := jb.Read16Bit()

	jb.MMU.Write(Address(pointer), byte(jb.CPU.ReadRegister(A)))

	return nil
}

func LDRAMPointerToA(jb *Jamboy, opcode OpCode) (err error) {
	pointer := jb.Read16Bit()

	jb.CPU.WriteRegister(A, uint(jb.MMU.Read(Address(pointer))))

	return nil
}

func LDH(jb *Jamboy, opcode OpCode) (err error) {
	offset := jb.Read8Bit()

	switch opcode & 0xF0 {
	case 0xE0:
		jb.MMU.Write(Address(0xFF00+uint(offset)), byte(jb.CPU.ReadRegister(A)))
	case 0xF0:
		jb.CPU.WriteRegister(A, uint(jb.MMU.Read(Address(0xFF00+uint(offset)))))
	}

	return nil
}

func LD(jb *Jamboy, opcode OpCode) (err error) {
	// 8 bit LDs
	if opcode >= 0x40 && opcode < 0x80 {
		opOffset := opcode - 0x40
		opSeq := opOffset % 8
		dstRegister := fullOrderedRegisters[opOffset/8]
		srcRegister := fullOrderedRegisters[opSeq]

		if dstRegister == HL {
			// (HL) = r
			jb.MMU.Write(Address(jb.CPU.ReadRegister(HL)), byte(jb.CPU.ReadRegisterInstant(srcRegister)))
		} else if opSeq == 6 {
			// r = (HL)
			jb.CPU.WriteRegisterInstant(dstRegister, uint(jb.MMU.Read(Address(jb.CPU.ReadRegister(HL)))))
		} else {
			//r = r
			jb.CPU.WriteRegister(dstRegister, uint(jb.CPU.ReadRegisterInstant(srcRegister)))
		}
	} else {
		panic(fmt.Sprintf("not implemented LD %x", opcode))
	}

	return err
}

var ld8AOrderedRegistersA = []RegisterID{
	B, D, H, HL,
}

func LDd8A(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := ld8AOrderedRegistersA[opcode&0xF0>>4]

	if dstRegister == HL {
		jb.MMU.Write(Address(jb.CPU.ReadRegister(HL)), jb.Read8Bit())
	} else {
		jb.CPU.WriteRegister(dstRegister, uint(jb.Read8Bit()))
	}

	return err
}

var ld8OrderedRegistersB = []RegisterID{
	C, E, L, A,
}

func LDd8B(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := ld8OrderedRegistersB[opcode&0xF0>>4]

	jb.CPU.WriteRegister(dstRegister, uint(jb.Read8Bit()))

	return err
}

func LDd16(jb *Jamboy, opcode OpCode) (err error) {
	register := RegisterID(0)

	switch opcode {
	case 0x01:
		register = BC
	case 0x11:
		register = DE
	case 0x21:
		register = HL
	case 0x31:
		register = SP
	default:
		panic(fmt.Sprintf("not implemented op LD -  %x", opcode))
	}

	jb.CPU.WriteRegister(register, uint(jb.Read16Bit()))

	return err
}

func LDRAMToA(jb *Jamboy, opcode OpCode) (err error) {
	data := byte(0)

	switch opcode & 0xF0 {
	case 0x00:
		data = jb.MMU.Read(Address(jb.CPU.ReadRegister(BC)))
	case 0x10:
		data = jb.MMU.Read(Address(jb.CPU.ReadRegister(DE)))
	case 0x20:
		data = jb.MMU.Read(Address(jb.CPU.ReadRegister(HL)))
		jb.CPU.IncrementHL()
	case 0x30:
		jb.CPU.DecrementHL()
		data = jb.MMU.Read(Address(jb.CPU.ReadRegister(HL)))
	}

	jb.CPU.WriteRegister(A, uint(data))

	return err
}

func LDAToRAM(jb *Jamboy, opcode OpCode) (err error) {
	data := jb.CPU.Registers[A]

	switch opcode & 0xF0 {
	case 0x00:
		register := jb.CPU.ReadRegister(BC)
		jb.MMU.Write(Address(register), data)
	case 0x10:
		register := jb.CPU.ReadRegister(DE)
		jb.MMU.Write(Address(register), data)
	case 0x20:
		register := jb.CPU.ReadRegister(HL)
		jb.MMU.Write(Address(register), data)
		jb.CPU.IncrementHL()
	case 0x30:
		jb.CPU.DecrementHL()
		register := jb.CPU.ReadRegister(HL)
		jb.MMU.Write(Address(register), data)
	}

	return err
}

func LDAToHighRAM(jb *Jamboy, opcode OpCode) (err error) {
	offset := jb.Read8Bit()

	jb.MMU.Write(Address(0xFF00+uint(offset)), byte(jb.CPU.ReadRegisterInstant(A)))

	return nil
}

func LDHighRAMToA(jb *Jamboy, opcode OpCode) (err error) {
	offset := jb.Read8Bit()

	jb.CPU.WriteRegister(A, uint(jb.MMU.ReadInstant(Address(0xFF00+uint(offset)))))

	return nil
}
