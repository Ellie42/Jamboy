package engine

import (
	"git.agehadev.com/elliebelly/gooey/lib/binary"
	"git.agehadev.com/elliebelly/jamboy/internal/code"
)

type Code struct {
	br       *binary.ByteReader
	opBuffer []code.Op
	jamboy   *Jamboy
	pc       int
	lc       int
	opMod    int
}

func (c *Code) NextOp() {
}

func (c *Code) Reset() {

}

func (c *Code) GetOpAtLine(index int) *code.Op {
	next := c.nextByte()

	op := code.Ops[int(next)+c.opMod]

	if op.Type == code.OpPREFIX {
		c.opMod = 0x0100
	}

	if op.Operands != nil {
		for _, operand := range op.Operands {
			if operand.ValueType != code.ValTypeRead {
				continue
			}

			if operand.ValueSize == 1 {
				val := c.nextByte()
				operand.ValueEvaluated = int(val)
			}
		}
	}

	c.opMod = 0

	c.opBuffer[index] = op

	return &op
}

func (c *Code) nextByte() byte {
	next := c.jamboy.MMU.ReadInstant(Address(c.pc))
	c.pc += 1
	return next
}

func NewCode(jamboy *Jamboy) *Code {
	return &Code{
		jamboy: jamboy,
	}
}
