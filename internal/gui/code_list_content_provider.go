package gui

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/behaviour"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
)

type CodeListContentWidget struct {
	widget.LinearLayout
	widgetListItem          *widget.WidgetListItem
	ByteNumberTextWidget    *widget.Text
	CommandStringTextWidget *widget.Text
}

func NewCodeListContentWidget(wli *widget.WidgetListItem) *CodeListContentWidget {
	cl := &CodeListContentWidget{
		widgetListItem: wli,
	}

	cl.FitToChildren = true

	cl.Rect = dimension.Rect{0, 0, 1, 1}

	leftPanelStyle := settings.WidgetPreferences{
		Name: fmt.Sprintf("Byte Number %d", wli.Index),
		Padding: &dimension.DirectionalRectSized{
			Left: dimension.Size{4, dimension.SizeUnitPixels},
		},
		DimensionBounds: &dimension.Dimensions{
			Width: &dimension.Size{
				Amount: 65,
			},
		},
		StyleSettings: &styles.StyleSettings{
			BackgroundColour: &ColourDarkRed,
		},
		Behaviours: &behaviour.BehaviourSet{
			Clickable: &behaviour.Clickable{},
		},
	}

	rightPanelStyle := settings.WidgetPreferences{
		Name:                fmt.Sprintf("Code Line %d", wli.Index),
		AlignmentHorizontal: settings.HorizontalLeft,
		Padding: &dimension.DirectionalRectSized{
			Left: dimension.Size{4, dimension.SizeUnitPixels},
		},
	}

	cl.ByteNumberTextWidget = widget.NewTextWidget(nil, "")
	cl.CommandStringTextWidget = widget.NewTextWidget(&settings.WidgetPreferences{
		AlignmentHorizontal: settings.HorizontalLeft,
	}, "")

	cl.AddChildWithParent(cl,
		widget.NewPanel(&leftPanelStyle, cl.ByteNumberTextWidget),
		widget.NewPanel(&rightPanelStyle, cl.CommandStringTextWidget),
	)

	return cl
}
