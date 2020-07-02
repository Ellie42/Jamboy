package engine

var orderedExtraOps = []RegisterID{
	B, C, D, E, H, L, HL, A, B, C, D, E, H, L, HL, A,
}

func getRegisterValueFromOpcode(jb *Jamboy, opcode OpCode) byte {
	register := orderedExtraOps[opcode&0x0F]

	if register == HL {
		return jb.MMU.Read(Address(jb.CPU.ReadRegisterInstant(register)))
	}

	return byte(jb.CPU.ReadRegister(register))
}

func BIT(jb *Jamboy, opcode OpCode) error {
	bitToTest := byte((opcode - 0x40) / 8)
	testByte := getRegisterValueFromOpcode(jb, opcode) & (1 << bitToTest)

	jb.CPU.SetFlagsMasked(
		SubFlag|HalfCarryFlag|(Flag(((^(testByte >> bitToTest))&1)*uint8(ZeroFlag))),
		ZeroFlag|HalfCarryFlag|SubFlag,
	)

	return nil
}

func RES(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op RES")
}

func RL(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op RL")
}

func RLC(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op RLC")
}

func RR(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op RR")
}

func RRC(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op RRC")
}

func SET(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op SET")
}

func SLA(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op SLA")
}

func SRA(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op SRA")
}

func SRL(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op SRL")
}

func SWAP(jb *Jamboy, opcode OpCode) error {
	panic("not implemented op SWAP")
}
