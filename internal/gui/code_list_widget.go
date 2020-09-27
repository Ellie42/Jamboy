package gui

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/jamboy/internal/code"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"strings"
)

type CodeListDataProvider struct {
	jamboy *engine.Jamboy
}

type CodeListData struct {
	byteNumber  int
	commandText string
}

var byteBuffer = make([]byte, 2)

func (c CodeListDataProvider) Provide(index int) CodeListData {
	op := c.jamboy.Code.GetOpAtLine(index)

	commandString := op.Type.String()

	commandString += strings.Repeat(" ", 4-len(commandString))

	operandStrings := make([]string, 0)

	if op.Operands != nil {
		for _, operand := range op.Operands {
			operandString := code.GetValueTypeString(operand.ValueStatic, operand.ValueType)

			if operand.ValueType == code.ValTypeRead {
				binary.BigEndian.PutUint16(byteBuffer, uint16(operand.ValueEvaluated))
				operandString = fmt.Sprintf("$%s", hex.EncodeToString(byteBuffer))
			} else if operand.ValueType == code.ValTypeRegister {
				operandString = engine.RegisterID((operand.ValueStatic)).String()
			} else if operand.ValueType == code.ValTypeKeyword {
				operandString = engine.Keyword((operand.ValueStatic)).String()
			}

			if operand.IncDecModifier > 0 {
				operandString += "+"
			} else if operand.IncDecModifier < 0 {
				operandString += "-"
			}

			if operand.RetrieveType == code.RetrievePointer {
				operandString = fmt.Sprintf("(%s)", operandString)
			}

			operandStrings = append(operandStrings, operandString)
		}

		// For some reason I dont feel like looking into right now the previous characters are not being cleared
		// properly, so I'm just wiping them clean with a bunch of spaces
		// TODO like fix it
		commandString += fmt.Sprintf(" %s                        ", strings.Join(operandStrings, ", "))
	}

	return CodeListData{
		op.ByteOffset,
		commandString,
	}
}

func (c *CodeListDataProvider) Init(jamboy *engine.Jamboy) {
	c.jamboy = jamboy
}

func getCodeListContentWidget(wli *widget.WidgetListItem) widget.Widget {
	return NewCodeListContentWidget(wli)
}

func NewCodeListWidget(jamboy *engine.Jamboy, debugger *DebuggerControl) *widget.List {
	dataProvider := &CodeListDataProvider{}

	dataProvider.Init(jamboy)

	list := widget.NewList(getCodeListContentWidget, func(w widget.Widget, index int) {
		t := w.(*CodeListContentWidget)
		data := dataProvider.Provide(index + jamboy.Debugger.GetCurrentLine())

		t.ByteNumberTextWidget.Text = fmt.Sprintf("%04x", data.byteNumber)

		t.CommandStringTextWidget.Text = data.commandText
	}, nil)

	return list
}
