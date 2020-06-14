package engine

import (
    "fmt"
    "git.agehadev.com/elliebelly/jamboy/internal"
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

func INC(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op INC -  %x", opcode))
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

var ldOrderedRegisters = []Register{
    B, C, D, E, H, L, Register(255), A,
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

func LD(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    // 8 bit LDs r = r
    if opcode >= 0x40 && opcode < 0x80 {
        opOffset := opcode - 0x40
        opSeq := opOffset % 8
        dstRegister := ldOrderedRegisters[opOffset/8]
        srcRegister := ldOrderedRegisters[opSeq]

        if dstRegister == 255 {
            // (HL) = r
            jb.Memory.RAM[jb.CPU.ReadRegister(HL)] = jb.CPU.Registers[srcRegister]
        }else if opSeq == 6 {
            // r = (HL)
            jb.CPU.Registers[dstRegister] = jb.Memory.RAM[jb.CPU.ReadRegister(HL)]
        } else {
            jb.CPU.WriteRegister(dstRegister, uint(jb.CPU.Registers[srcRegister]))
        }
    }

    return true, err
}

func LDH(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    panic(fmt.Sprintf("not implemented op LDH -  %x", opcode))
}

func NOP(jb *Jamboy, opcode OpCode) (finished bool, err error) {
    internal.Logger.Info("exec NOP")

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
