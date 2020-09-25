package renderer

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"math/rand"
)

type CodeListDataProvider struct {
	jamboy *engine.Jamboy
}

type CodeListData struct {
	byteNumber  int
	commandText string
}

func (c CodeListDataProvider) Provide(index int) CodeListData {
	c.jamboy.Code.GetOpAtLine(index)

	return CodeListData{
		index * 2,
		"LD A, B",
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
			t.CommandStringTextWidget.Text = fmt.Sprintf("LD AF, %04x", rand.Intn(0xFFFF))
		}
	}, nil)

	return list
}
