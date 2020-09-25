package engine

import (
	"git.agehadev.com/elliebelly/gooey/lib/binary"
	"git.agehadev.com/elliebelly/jamboy/internal/code"
)

type Code struct {
	br       *binary.ByteReader
	opBuffer []code.Op
}

func (c *Code) NextOp() {
}

func (c *Code) Reset() {

}

func (c *Code) GetOpAtLine(index int) {

}

func NewCode(*Jamboy) *Code {
	return &Code{}
}
