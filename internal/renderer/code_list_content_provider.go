package renderer

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
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
		Padding: &dimension.DirectionalRectSized{
			Left: dimension.Size{0.1, dimension.SizeUnitPixels},
		},
		DimensionBounds: &dimension.Dimensions{
			Width: &dimension.Size{
				Amount: 65,
			},
		},
		StyleSettings: &styles.StyleSettings{
			BackgroundColour: &ColourBGGreyLight,
		},
	}

	rightPanelStyle := settings.WidgetPreferences{
		AlignmentHorizontal: settings.HorizontalLeft,
		Padding: &dimension.DirectionalRectSized{
			Left: dimension.Size{32, dimension.SizeUnitPixels},
		},
		//StyleSettings: &styles.StyleSettings{
		//	BackgroundColour: &ColourBGPink,
		//},
	}

	cl.ByteNumberTextWidget = widget.NewTextWidget(nil)
	cl.CommandStringTextWidget = widget.NewTextWidget(&settings.WidgetPreferences{
		AlignmentHorizontal: settings.HorizontalLeft,
	})

	cl.AddChildWithParent(cl,
		widget.NewPanel(&leftPanelStyle, cl.ByteNumberTextWidget),
		widget.NewPanel(&rightPanelStyle, cl.CommandStringTextWidget),
	)

	return cl
}
