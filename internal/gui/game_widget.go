package gui

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"github.com/go-gl/gl/v4.6-compatibility/gl"
)

type GameWidget struct {
	widget.BaseWidget
	Game *Game
}

func (g *GameWidget) Init() {
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

	gl.Viewport(
		int32(0),
		int32(0),
		int32(widget.Context.Resolution.Width),
		int32(widget.Context.Resolution.Height),
	)
}

func calculateViewport(index int, parent widget.Widget) dimension.Rect {
	parentRect := parent.GetChildRectAbsolute(index)
	resolution := widget.Context.Resolution

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
