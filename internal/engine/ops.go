package engine

import (
    "fmt"
)

var fullOrderedRegisters = []RegisterID{
    B, C, D, E, H, L, RegisterID(255), A,
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
    panic(fmt.Sprintf("not implemented op CPL -  %x", opcode))
}

func DAA(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op DAA -  %x", opcode))
}

func DEC(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op DEC -  %x", opcode))
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
    panic(fmt.Sprintf("not implemented op HALT -  %x", opcode))
}

var incdecAOrderedRegisters = []RegisterID{
    B, D, H, HL,
}

var incdecBOrderedRegisters = []RegisterID{
    C, E, L, A,
}

func INCA(jb *Jamboy, opcode OpCode) (err error) {
    dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]
    IncrementRegister(jb, dstRegister)

    return nil
}

func IncrementRegister(jb *Jamboy, dstRegister RegisterID) {
    jb.CPU.Add(dstRegister, 1, false)
}

func DecrementRegister(jb *Jamboy, dstRegister RegisterID) {
    jb.CPU.Subtract(dstRegister, 1, false)
}

func INCB(jb *Jamboy, opcode OpCode) (err error) {
    dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

    IncrementRegister(jb, dstRegister)

    return nil
}

func DECA(jb *Jamboy, opcode OpCode) (err error) {
    dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]

    DecrementRegister(jb, dstRegister)

    return nil
}

func DECB(jb *Jamboy, opcode OpCode) (err error) {
    dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

    DecrementRegister(jb, dstRegister)

    return nil
}

func JP(jb *Jamboy, opcode OpCode) (err error) {
    address := jb.Read16Bit()

    jb.CPU.WriteRegister(PC, uint(address))

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
    panic(fmt.Sprintf("not implemented op PREFIX -  %x", opcode))
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
    panic(fmt.Sprintf("not implemented op RLA -  %x", opcode))
}

func RLCA(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op RLCA -  %x", opcode))
}

func RRA(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op RRA -  %x", opcode))
}

func RRCA(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op RRCA -  %x", opcode))
}

func RST(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op RST -  %x", opcode))
}

func SCF(jb *Jamboy, opcode OpCode) (err error) {
    jb.CPU.SetFlags(CarryFlag | (jb.CPU.GetFlags() & 0x80))

    return nil
}

func STOP(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op STOP -  %x", opcode))
}
