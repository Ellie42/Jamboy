package engine

import (
    "fmt"
)

var fullOrderedRegisters = []RegisterID{
    B, C, D, E, H, L, RegisterID(255), A,
}

func ADC(jb *Jamboy, opcode OpCode) (err error) {
    register := fullOrderedRegisters[opcode&0x0F]

    jb.CPU.AddR8(A, register, true)

    return nil
}

func ADD(jb *Jamboy, opcode OpCode) (err error) {
    register := fullOrderedRegisters[opcode&0x0F]

    jb.CPU.AddR8(A, register, false)

    return nil
}

func AND(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op AND -  %x", opcode))
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
    panic(fmt.Sprintf("not implemented op CP -  %x", opcode))
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

var ld8OrderedRegisters = []RegisterID{
    C, E, L, A,
}

func LDd8(jb *Jamboy, opcode OpCode) (err error) {
    dstRegister := ld8OrderedRegisters[opcode&0xF0>>4]

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

func LD(jb *Jamboy, opcode OpCode) (err error) {
    // 8 bit LDs
    if opcode >= 0x40 && opcode < 0x80 {
        opOffset := opcode - 0x40
        opSeq := opOffset % 8
        dstRegister := fullOrderedRegisters[opOffset/8]
        srcRegister := fullOrderedRegisters[opSeq]

        if dstRegister == 255 {
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

func NOP(jb *Jamboy, opcode OpCode) (err error) {
    return nil
}

func OR(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op OR -  %x", opcode))
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

func SBC(jb *Jamboy, opcode OpCode) (err error) {
    register := fullOrderedRegisters[opcode&0x0F]

    jb.CPU.SubtractR8(A, register, true)

    return nil
}

func SCF(jb *Jamboy, opcode OpCode) (err error) {
    jb.CPU.SetFlags(CarryFlag | (jb.CPU.GetFlags() & 0x80))

    return nil
}

func STOP(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op STOP -  %x", opcode))
}

func SUB(jb *Jamboy, opcode OpCode) (err error) {
    register := fullOrderedRegisters[opcode&0x0F]

    jb.CPU.SubtractR8(A, register, false)

    return nil
}

func XOR(jb *Jamboy, opcode OpCode) (err error) {
    panic(fmt.Sprintf("not implemented op XOR -  %x", opcode))
}
