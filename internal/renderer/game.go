package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"unsafe"
)

type Game struct {
	Pixels                  []uint8
	textureHandle           uint32
	quadVao                 uint32
	UVBO                    uint32
	PanelWidth              float32
	gameBoyAspectRatioMulti float32

	pixelsPointer unsafe.Pointer
	ResX          int32
	ResY          int32
}

func (g Game) SetupGamePanel(w *Window) {
	resizedQuad := make([]float32, len(quad))

	copy(resizedQuad, quad)

	width, height := w.glfwWindow.GetSize()
	windowAspect := float32(width) / float32(height)
	panelWidth := w.Game.PanelWidth * 2

	for i := 0; i < len(resizedQuad); i += 3 {
		x := resizedQuad[i]
		y := resizedQuad[i+1]

		resizedQuad[i] = -1 + x*panelWidth

		windowHeight := panelWidth * w.Game.gameBoyAspectRatioMulti * windowAspect

		resizedQuad[i+1] = (1 - windowHeight) + y*windowHeight
	}

	w.Game.quadVao, w.Game.UVBO = makeVao(resizedQuad, quadUVs)
}

func (g Game) Render(w *Window) {
	gl.BindVertexArray(w.Game.quadVao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, w.Game.textureHandle)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		g.ResX,
		g.ResY,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		g.pixelsPointer,
	)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quad)/3))
}

func (g Game) InitGL(w *Window) {
	gl.GenTextures(1, &w.Game.textureHandle)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, w.Game.textureHandle)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, 0)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, 0)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.NEAREST)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	g.SetupGamePanel(w)
}

func (g Game) OnWindowResize(w *Window) {
	g.SetupGamePanel(w)
}

var quad = []float32{
	0, 1, 0, // Top Left
	0, 0, 0, // Bottom Left
	1, 0, 0, // Bottom Right

	0, 1, 0, // Top Left
	1, 0, 0, // Bottom Right
	1, 1, 0, // Top Right
}

var quadUVs = []float32{
	0, 0, // Top Left
	0, 1, // Bottom Left
	1, 1, // Bottom Right

	0, 0, // Top Left
	1, 1, // Bottom Right
	1, 0, // Top Right
}
