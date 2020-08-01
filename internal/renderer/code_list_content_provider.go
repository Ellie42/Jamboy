package renderer

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
)

type CodeListContentWidgetConfig struct {
	widget                  widget.Widget
	byteNumberTextWidget    *widget.Text
	commandStringTextWidget *widget.Text
	byteNumber              int
}

type CodeListContentProvider struct {
	widgets            map[int]*CodeListContentWidgetConfig
	WidgetDataProvider *CodeListDataProvider
}

func (c *CodeListContentProvider) InitListItem(w *widget.WidgetListItem) {
	index := w.GetIndex()

	if c.widgets == nil {
		c.widgets = make(map[int]*CodeListContentWidgetConfig)
	}

	var ok bool

	if _, ok = c.widgets[index]; ok {
		return
	}

	leftPanelStyle := settings.WidgetPreferences{
		Padding: &dimension.DirectionalRect{
			Left:   0.1,
		},
		DimensionBounds: &dimension.Dimensions{
			Width: &dimension.Size{
				Amount: 100,
				Unit:   dimension.SizeUnitPixels,
			},
		},
		StyleSettings: &styles.StyleSettings{
			BackgroundColour: &ColourBGGrey,
		},
	}

	byteNumberTextWidget := widget.NewTextWidget(nil)
	commandStringTextWidget := widget.NewTextWidget(nil)

	c.widgets[index] = &CodeListContentWidgetConfig{
		widget: widget.NewLinearLayout(
			nil,
			widget.NewPanel(&leftPanelStyle, byteNumberTextWidget),
			widget.NewPanel(nil, commandStringTextWidget),
		),
		byteNumberTextWidget:    byteNumberTextWidget,
		commandStringTextWidget: commandStringTextWidget,
	}

	c.widgets[index].widget.SetParent(w)
	c.widgets[index].widget.Init()
}

func (c *CodeListContentProvider) RenderListItem(w *widget.WidgetListItem) {
	wConfig := c.widgets[w.GetIndex()]
	data := c.WidgetDataProvider.Provide(w.GetIndex())

	wConfig.byteNumberTextWidget.Text = fmt.Sprintf("0x%04x", data.byteNumber)
	wConfig.commandStringTextWidget.Text = data.commandText

	c.widgets[w.GetIndex()].widget.Render()
}

func NewCodeListContentProvider(dataProvider *CodeListDataProvider) *CodeListContentProvider {
	c := &CodeListContentProvider{
		WidgetDataProvider: dataProvider,
	}

	return c
}
