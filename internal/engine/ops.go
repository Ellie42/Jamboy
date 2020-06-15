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

var incOrderedRegisters = []Register{
    C, E, L, A,
}

func INC(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := incOrderedRegisters[(opcode&0xF0)>>4]

    jb.CPU.WriteRegister(dstRegister, jb.CPU.ReadRegister(dstRegister)+1)

    return true, nil
}

func JP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    address := jb.Read16Bit()
    jb.CPU.Wait(1)
    jb.CPU.PC = address

    return true, nil
}

func JR(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op JR -  %x", opcode))
}

func LDRAMToA(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    data := byte(0)

    switch opcode & 0xF0 {
    case 0x00:
        data = jb.Memory.RAM[jb.CPU.ReadRegister(BC)]
    case 0x10:
        data = jb.Memory.RAM[jb.CPU.ReadRegister(DE)]
    case 0x20:
        data = jb.Memory.RAM[jb.CPU.ReadRegister(HL)]
        jb.CPU.IncrementHL()
    case 0x30:
        data = jb.Memory.RAM[jb.CPU.ReadRegister(HL)]
        jb.CPU.DecrementHL()
    }

    jb.CPU.WriteRegister(A, uint(data))

    return true, err
}

func LDAToRAM(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    data := jb.CPU.Registers[A]

    switch opcode & 0xF0 {
    case 0x00:
        register := jb.CPU.ReadRegister(BC)
        jb.CPU.Wait(1)
        jb.Memory.Write(uint16(register), data)
    case 0x10:
        register := jb.CPU.ReadRegister(DE)
        jb.CPU.Wait(1)
        jb.Memory.Write(uint16(register), data)
    case 0x20:
        register := jb.CPU.ReadRegister(HL)
        jb.CPU.Wait(1)
        jb.Memory.Write(uint16(register), data)
        jb.CPU.IncrementHL()
    case 0x30:
        register := jb.CPU.ReadRegister(HL)
        jb.CPU.Wait(1)
        jb.Memory.Write(uint16(register), data)
        jb.CPU.DecrementHL()
    }

    return true, err
}

var ld8OrderedRegisters = []Register{
    C, E, L, A,
}

func LDd8(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    dstRegister := ld8OrderedRegisters[opcode&0xF0>>4]

    jb.CPU.WriteRegister(dstRegister, uint(jb.Read8Bit()))

    return true, err
}

func LDd16(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    register := Register(0)

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

var ldOrderedRegisters = []Register{
    B, C, D, E, H, L, Register(255), A,
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
            jb.Memory.RAM[jb.CPU.ReadRegister(HL)] = jb.CPU.Registers[srcRegister]
        } else if opSeq == 6 {
            // r = (HL)
            jb.CPU.Registers[dstRegister] = jb.Memory.RAM[jb.CPU.ReadRegister(HL)]
        } else {
            //r = r
            jb.CPU.WriteRegister(dstRegister, uint(jb.CPU.Registers[srcRegister]))
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
