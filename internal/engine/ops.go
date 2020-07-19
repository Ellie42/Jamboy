package engine

import (
	"fmt"
)

var fullOrderedRegisters = []RegisterID{
	B, C, D, E, H, L, HL, A,
}

func CALL(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.Wait(1)

	switch opcode {
	case 0xC4:
		if !jb.CPU.IsFlagSet(ZeroFlag) {
			jb.CPU.Call(jb.Read16Bit())
		}
	case 0xD4:
		if !jb.CPU.IsFlagSet(CarryFlag) {
			jb.CPU.Call(jb.Read16Bit())
		}
	case 0xCC:
		if jb.CPU.IsFlagSet(ZeroFlag) {
			jb.CPU.Call(jb.Read16Bit())
		}
	case 0xDC:
		if jb.CPU.IsFlagSet(CarryFlag) {
			jb.CPU.Call(jb.Read16Bit())
		}
	case 0xCD:
		jb.CPU.Call(jb.Read16Bit())
	default:
		panic(fmt.Sprintf("not implemented CALL: %x", opcode))
	}

	return nil
}

func CCF(jb *Jamboy, opcode OpCode) (err error) {
	panic(fmt.Sprintf("not implemented op CCF -  %x", opcode))
}

func CP(jb *Jamboy, opcode OpCode) (err error) {
	register := getOrderedRegister(opcode)

	if opcode&0xF0 == 0xF0 {
		jb.CPU.Compare(jb.Read8Bit())
		return nil
	}

	jb.CPU.CompareR8(register)

	return nil
}

func CPL(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.WriteRegister(A, ^jb.CPU.ReadRegisterInstant(A))

	jb.CPU.AddFlags(SubFlag | HalfCarryFlag)

	return nil
}

func DAA(jb *Jamboy, opcode OpCode) (err error) {
	panic(fmt.Sprintf("not implemented op DAA -  %x", opcode))
}

var incdec16OrderedRegisters = []RegisterID{
	BC, DE, HL, SP,
}

func DEC16(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.Decrement(incdec16OrderedRegisters[opcode&0xF0>>4])
	return nil
}

func DI(jb *Jamboy, opcode OpCode) (err error) {
	jb.MMU.Write(0xFFFF, 0)

	return nil
}

func EI(jb *Jamboy, opcode OpCode) (err error) {
	jb.MMU.Write(0xFFFF, 1)

	return nil
}

func HALT(jb *Jamboy, opcode OpCode) (err error) {
	jb.Halt()
	return nil
}

func INC16(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := incdec16OrderedRegisters[(opcode&0xF0)>>4]
	jb.CPU.Increment(dstRegister)

	return nil
}

var incdecAOrderedRegisters = []RegisterID{
	B, D, H, HL,
}

var incdecBOrderedRegisters = []RegisterID{
	C, E, L, A,
}

func INC_0(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]
	IncrementRegister(jb, dstRegister)

	return nil
}

func IncrementRegister(jb *Jamboy, dstRegister RegisterID) {
	jb.CPU.Add(dstRegister, 1, false, false)
}

func DecrementRegister(jb *Jamboy, dstRegister RegisterID) {
	jb.CPU.Subtract(dstRegister, 1, false, false)
}

func INC_1(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

	IncrementRegister(jb, dstRegister)

	return nil
}

func DEC_0(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]

	DecrementRegister(jb, dstRegister)

	return nil
}

func DEC_1(jb *Jamboy, opcode OpCode) (err error) {
	dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

	DecrementRegister(jb, dstRegister)

	return nil
}

func JP(jb *Jamboy, opcode OpCode) (err error) {
	address := jb.Read16Bit()

	jb.CPU.Jump(uint16(address))

	return nil
}

func JR(jb *Jamboy, opcode OpCode) (err error) {
	offset := jb.Read8Bit()

	switch opcode {
	case 0x18:
		jb.CPU.Jump(uint16(int(jb.CPU.ReadRegister(PC)) + int(int8(offset))))
	case 0x20:
		if !jb.CPU.IsFlagSet(ZeroFlag) {
			jb.CPU.Jump(uint16(int(jb.CPU.ReadRegister(PC)) + int(int8(offset))))
		}
	case 0x30:
		if !jb.CPU.IsFlagSet(CarryFlag) {
			jb.CPU.Jump(uint16(int(jb.CPU.ReadRegister(PC)) + int(int8(offset))))
		}
	case 0x28:
		if jb.CPU.IsFlagSet(ZeroFlag) {
			jb.CPU.Jump(uint16(int(jb.CPU.ReadRegister(PC)) + int(int8(offset))))
		}
	case 0x38:
		if jb.CPU.IsFlagSet(CarryFlag) {
			jb.CPU.Jump(uint16(int(jb.CPU.ReadRegister(PC)) + int(int8(offset))))
		}
	default:
		panic(fmt.Sprintf("not implemented JR %02x", opcode))
	}

	return err
}

func NOP(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.Wait(1)
	return nil
}

var pushPopOrderedRegisters = []RegisterID{
	BC, DE, HL, AF,
}

func POP(jb *Jamboy, opcode OpCode) (err error) {
	register := pushPopOrderedRegisters[(opcode&0xF0>>4)-0x0C]

	value := jb.CPU.PopUint16()
	jb.CPU.WriteRegister(register, uint(value))

	return nil
}

func PREFIX(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.CurrentJumpTable = &ExtraOpJumpTable

	return nil
}

func PUSH(jb *Jamboy, opcode OpCode) (err error) {
	register := pushPopOrderedRegisters[(opcode&0xF0>>4)-0x0C]

	value := jb.CPU.ReadRegister(register)
	jb.CPU.PushUint16(uint16(value))

	return nil
}

func RET(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.Return()

	return nil
}

func RETI(jb *Jamboy, opcode OpCode) (err error) {
	panic(fmt.Sprintf("not implemented op RETI -  %x", opcode))
}

func RLA(jb *Jamboy, opcode OpCode) (err error) {
	val := jb.CPU.ReadRegisterInstant(A)

	carry := jb.CPU.GetFlag(CarryFlag)

	rotateLeftASetCarryFlag(jb, val, carry)

	return nil
}

func rotateLeftASetCarryFlag(jb *Jamboy, val uint, rotateBit uint8) {
	jb.CPU.SetFlags(0x00)

	if val&0x80 > 0 {
		jb.CPU.AddFlags(CarryFlag)
	}

	val <<= 1
	val |= uint(rotateBit)

	jb.CPU.WriteRegister(A, val)
}

func RLCA(jb *Jamboy, opcode OpCode) (err error) {
	val := jb.CPU.ReadRegisterInstant(A)

	rotateLeftASetCarryFlag(jb, val, uint8(val&0x80)>>7)

	return nil
}

func RRA(jb *Jamboy, opcode OpCode) (err error) {
	val := jb.CPU.ReadRegisterInstant(A)

	carry := jb.CPU.GetFlag(CarryFlag)

	rotateRightASetCarryFlag(jb, val, carry)

	return nil
}

func rotateRightASetCarryFlag(jb *Jamboy, val uint, rotateBit uint8) {
	jb.CPU.SetFlags(0x00)

	if val&0x01 > 0 {
		jb.CPU.AddFlags(CarryFlag)
	}

	val >>= 1
	val |= uint(rotateBit << 7)

	jb.CPU.WriteRegister(A, val)
}

func RRCA(jb *Jamboy, opcode OpCode) (err error) {
	val := jb.CPU.ReadRegisterInstant(A)

	rotateRightASetCarryFlag(jb, val, uint8(val&1))

	return nil
}

func RST(jb *Jamboy, opcode OpCode) (err error) {
	addr := 0

	switch opcode {
	case 0xC7:
		addr = 0x00
	case 0xD7:
		addr = 0x10
	case 0xE7:
		addr = 0x20
	case 0xF7:
		addr = 0x30
	case 0xCF:
		addr = 0x08
	case 0xDF:
		addr = 0x18
	case 0xEF:
		addr = 0x28
	case 0xFF:
		addr = 0x38
	default:
		panic(fmt.Sprintf("unsupported RST op %x", opcode))
	}

	jb.CPU.Call(uint16(addr))

	return
}

func SCF(jb *Jamboy, opcode OpCode) (err error) {
	jb.CPU.SetFlags(CarryFlag | (jb.CPU.GetFlags() & 0x80))

	return nil
}

func STOP(jb *Jamboy, opcode OpCode) (err error) {
	panic(fmt.Sprintf("not implemented op STOP -  %x", opcode))
}
