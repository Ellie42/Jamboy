package engine

import (
	"git.agehadev.com/elliebelly/gooey/lib/binary"
	"git.agehadev.com/elliebelly/jamboy/internal/code"
)

//go:generate stringer -type Keyword -linecomment
type Keyword int

const (
	KeywordZ  Keyword = iota //Z
	KeywordNZ                //NZ
	KeywordCB                //CB
	KeywordC                 //C
	KeywordNC                //NC
)

type Code struct {
	br       *binary.ByteReader
	opBuffer []*code.Op
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
	if c.opBuffer[index] != nil {
		return c.opBuffer[index]
	}

	for i := c.lc; i <= index; i++ {
		c.parseNextLine(i)
	}

	return c.opBuffer[index]
}

func (c *Code) parseNextLine(index int) *code.Op {
	currentPC := c.pc

	next := c.nextByte()

	op := code.Ops[int(next)+c.opMod]

	op.ByteOffset = currentPC

	if op.Type == code.OpPREFIX {
		c.opMod = 0x0100
		return c.parseNextLine(index)
	}

	if op.Operands != nil {
		for i := 0; i < len(op.Operands); i++ {
			operand := &op.Operands[i]

			if operand.ValueType != code.ValTypeRead || operand.ValueSize == 0 {
				continue
			}

			valueLocation := code.ValueLocation(operand.ValueStatic)

			if operand.ValueSize == 1 {
				val := c.nextByte()

				if valueLocation == code.ValAbsolute {
					operand.ValueEvaluated = int(uint(val))
				} else {
					operand.ValueEvaluated = int(val)
				}
			} else if operand.ValueSize == 2 {
				valLow := c.nextByte()
				valHigh := c.nextByte()

				val := (uint16(valHigh) << 8) | (uint16(valLow))

				if valueLocation == code.ValAbsolute {
					operand.ValueEvaluated = int(uint(val))
				} else {
					//TODO verify this converts from uint16 to int correctly
					operand.ValueEvaluated = int(val)
				}
			}
		}
	}

	c.opMod = 0

	c.opBuffer[index] = &op

	c.lc++

	return &op
}

func (c *Code) nextByte() byte {
	next := c.jamboy.MMU.ReadInstant(Address(c.pc))
	c.pc++
	return next
}

func NewCode(jamboy *Jamboy) *Code {
	return &Code{
		opBuffer: make([]*code.Op, 65535),
		jamboy:   jamboy,
	}
}
