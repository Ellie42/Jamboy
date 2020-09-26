package renderer

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

			if operand.RetrieveType == code.RetrievePointer {
				operandString = fmt.Sprintf("(%s)", operandString)
			}

			if operand.IncDecModifier > 0 {
				operandString += "+"
			} else if operand.IncDecModifier < 0 {
				operandString += "-"
			}

			operandStrings = append(operandStrings, operandString)
		}

		commandString += fmt.Sprintf(" %s", strings.Join(operandStrings, ", "))
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

func NewCodeListWidget(jamboy *engine.Jamboy) *widget.List {
	dataProvider := &CodeListDataProvider{}

	dataProvider.Init(jamboy)

	list := widget.NewList(getCodeListContentWidget, func(w widget.Widget, index int) {
		t := w.(*CodeListContentWidget)
		data := dataProvider.Provide(index)

		t.ByteNumberTextWidget.Text = fmt.Sprintf("%04x", data.byteNumber)

		if t.CommandStringTextWidget.Text == "" {
			t.CommandStringTextWidget.Text = data.commandText
		}
	}, nil)

	return list
}
