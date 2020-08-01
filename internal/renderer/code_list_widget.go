package renderer

import (
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
)

type CodeListDataProvider struct {
	Cart *engine.Cart
}

type CodeListData struct {
	byteNumber  int
	commandText string
}

func (c CodeListDataProvider) Provide(index int) CodeListData {
	return CodeListData{
		index * 2,
		"LD A, B",
	}
}

func GetCodeListWidget() *widget.List {
	dataProvider := &CodeListDataProvider{}

	//contentProvider := widget.NewStringListContentProvider(dataProvider)
	contentProvider := NewCodeListContentProvider(dataProvider)

	list := widget.NewList(contentProvider, nil)

	return list
}
