package engine

import (
    "fmt"
)

func ADC(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op ADC -  %x", opcode))
}

func ADD(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op ADD -  %x", opcode))
}

func AND(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op AND -  %x", opcode))
}

func CALL(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op CALL -  %x", opcode))
}

func CCF(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op CCF -  %x", opcode))
}

func CP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op CP -  %x", opcode))
}

func CPL(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op CPL -  %x", opcode))
}

func DAA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op DAA -  %x", opcode))
}

func DEC(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op DEC -  %x", opcode))
}

func DI(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op DI -  %x", opcode))
}

func EI(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op EI -  %x", opcode))
}

func HALT(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op HALT -  %x", opcode))
}

var incdecAOrderedRegisters = []RegisterID{
    B, D, H, HL,
}

var incdecBOrderedRegisters = []RegisterID{
    C, E, L, A,
}

func INCA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]
    IncrementRegister(jb, dstRegister)

    return true, nil
}

func IncrementRegister(jb *Jamboy, dstRegister RegisterID) {
    srcValue := jb.CPU.ReadRegister(dstRegister)
    finalValue := uint(srcValue)

    jb.CPU.SetFlags(Flag(jb.CPU.Registers[F] & 0x10))

    if dstRegister == HL {
        hl := srcValue
        srcValue = uint(jb.ReadRAM(Address(hl)))
        finalValue = srcValue + 1
        jb.WriteRAM(Address(hl), byte(finalValue))
    } else {
        finalValue = srcValue + 1
        jb.CPU.WriteRegister(dstRegister, finalValue)
    }

    if finalValue >= 0x100 {
        jb.CPU.AddFlags(ZeroFlag)
    }

    if srcValue < 0x10 && finalValue >= 0x10 {
        jb.CPU.AddFlags(HalfCarryFlag)
    }

    if finalValue > 0xFF {
        jb.CPU.AddFlags(CarryFlag)
    }
}

func DecrementRegister(jb *Jamboy, dstRegister RegisterID) {
    srcValue := int(jb.CPU.ReadRegister(dstRegister))
    finalValue := int(srcValue)

    jb.CPU.SetFlags(Flag(jb.CPU.Registers[F] & 0x10) | SubFlag)

    if dstRegister == HL {
        hl := srcValue
        srcValue = int(jb.ReadRAM(Address(hl)))
        finalValue = srcValue - 1
        jb.WriteRAM(Address(hl), byte(finalValue))
    } else {
        finalValue = srcValue - 1
        jb.CPU.WriteRegister(dstRegister, uint(finalValue))
    }

    if finalValue == 0 {
        jb.CPU.AddFlags(ZeroFlag)
    }

    if srcValue >= 0x10 && finalValue < 0x10 {
        jb.CPU.AddFlags(HalfCarryFlag)
    }

    if finalValue < 0 {
        jb.CPU.AddFlags(CarryFlag)
    }
}

func INCB(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

    IncrementRegister(jb, dstRegister)

    return true, nil
}

func DECA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := incdecAOrderedRegisters[(opcode&0xF0)>>4]

    DecrementRegister(jb, dstRegister)

    return true, nil
}

func DECB(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := incdecBOrderedRegisters[(opcode&0xF0)>>4]

    DecrementRegister(jb, dstRegister)

    return true, nil
}

func JP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    address := jb.Read16Bit()

    jb.CPU.WriteRegister(PC, uint(address))


    return true, nil
}

func JR(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    offset := jb.Read8Bit()

    switch opcode {
    case 0x20:
        if jb.CPU.GetFlags()&ZeroFlag > 0 {
            return
        }
    default:
        panic(fmt.Sprintf("not implemented JR %02x", opcode))
    }

    jb.CPU.WriteRegister(PC, uint(int(jb.CPU.ReadRegister(PC))+int(int8(offset))))

    return true, err
}

func LDRAMToA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    data := byte(0)

    switch opcode & 0xF0 {
    case 0x00:
        data = jb.ReadRAM(Address(jb.CPU.ReadRegister(BC)))
    case 0x10:
        data = jb.ReadRAM(Address(jb.CPU.ReadRegister(DE)))
    case 0x20:
        data = jb.ReadRAM(Address(jb.CPU.ReadRegister(HL)))
        jb.CPU.IncrementHL()
    case 0x30:
        jb.CPU.DecrementHL()
        data = jb.ReadRAM(Address(jb.CPU.ReadRegister(HL)))
    }

    jb.CPU.WriteRegister(A, uint(data))

    return true, err
}

func LDAToRAM(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    data := jb.CPU.Registers[A]

    switch opcode & 0xF0 {
    case 0x00:
        register := jb.CPU.ReadRegister(BC)
        jb.WriteRAM(Address(register), data)
    case 0x10:
        register := jb.CPU.ReadRegister(DE)
        jb.WriteRAM(Address(register), data)
    case 0x20:
        register := jb.CPU.ReadRegister(HL)
        jb.WriteRAM(Address(register), data)
        jb.CPU.IncrementHL()
    case 0x30:
        jb.CPU.DecrementHL()
        register := jb.CPU.ReadRegister(HL)
        jb.WriteRAM(Address(register), data)
    }

    return true, err
}

var ld8OrderedRegisters = []RegisterID{
    C, E, L, A,
}

func LDd8(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := ld8OrderedRegisters[opcode&0xF0>>4]

    jb.CPU.WriteRegister(dstRegister, uint(jb.Read8Bit()))

    return true, err
}

func LDd16(jb *Jamboy, opcode OpCode) (finished bool, err error) {
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

    return true, err
}

var ldOrderedRegisters = []RegisterID{
    B, C, D, E, H, L, RegisterID(255), A,
}

func LD(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    // 8 bit LDs
    if opcode >= 0x40 && opcode < 0x80 {
        opOffset := opcode - 0x40
        opSeq := opOffset % 8
        dstRegister := ldOrderedRegisters[opOffset/8]
        srcRegister := ldOrderedRegisters[opSeq]

        if dstRegister == 255 {
            // (HL) = r
            jb.WriteRAM(Address(jb.CPU.ReadRegister(HL)), byte(jb.CPU.ReadRegisterInstant(srcRegister)))
        } else if opSeq == 6 {
            // r = (HL)
            jb.CPU.WriteRegisterInstant(dstRegister, uint(jb.ReadRAM(Address(jb.CPU.ReadRegister(HL)))))
        } else {
            //r = r
            jb.CPU.WriteRegister(dstRegister, uint(jb.CPU.ReadRegisterInstant(srcRegister)))
        }
    } else {
        panic(fmt.Sprintf("not implemented LD %x", opcode))
    }

    return true, err
}

func LDH(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op LDH -  %x", opcode))
}

func NOP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    return true, nil
}

func OR(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op OR -  %x", opcode))
}

func POP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op POP -  %x", opcode))
}

func PREFIX(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op PREFIX -  %x", opcode))
}

func PUSH(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op PUSH -  %x", opcode))
}

func RET(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RET -  %x", opcode))
}

func RETI(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RETI -  %x", opcode))
}

func RLA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RLA -  %x", opcode))
}

func RLCA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RLCA -  %x", opcode))
}

func RRA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RRA -  %x", opcode))
}

func RRCA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RRCA -  %x", opcode))
}

func RST(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op RST -  %x", opcode))
}

func SBC(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op SBC -  %x", opcode))
}

func SCF(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op SCF -  %x", opcode))
}

func STOP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op STOP -  %x", opcode))
}

func SUB(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op SUB -  %x", opcode))
}

func XOR(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op XOR -  %x", opcode))
}
