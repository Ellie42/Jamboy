package renderer

import (
	"git.agehadev.com/elliebelly/gooey/pkg/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/window"
	"git.agehadev.com/elliebelly/gooey/pkg/window/widget"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type GameWidget struct {
	widget.BaseWidget
	Game *Game
}

func (g *GameWidget) Init(parent widget.WidgetParent) {
	g.Parent = parent
	g.Game.InitGL()
}

func (g GameWidget) Render() {
	viewportRect := calculateViewport(g.Index, g.Parent)

	gl.Viewport(
		int32(viewportRect.X),
		int32(viewportRect.Y),
		int32(viewportRect.Width),
		int32(viewportRect.Height),
	)

	g.Game.Render()
}

func calculateViewport(index int, parent widget.WidgetParent) dimension.Rect {
	parentRect := parent.GetChildRectRelative(index)
	resolution := window.Context.Resolution

	return dimension.Rect{
		X:      parentRect.X * float32(resolution.Width),
		Y:      parentRect.Y * float32(resolution.Height),
		Width:  parentRect.Width * float32(resolution.Width),
		Height: parentRect.Height * float32(resolution.Height),
	}
}

func NewGameWidget(game *Game) *GameWidget {
	gw := &GameWidget{
		Game: game,
	}

	return gw
}
