package renderer

import (
	"git.agehadev.com/elliebelly/gooey/lib/shader"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"github.com/AllenDang/giu"
	"github.com/go-gl/gl/v4.6-core/gl"
	"image"
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
	PixelTexure   *giu.Texture
	PixelImage    *image.RGBA
	ImageWidget   *giu.ImageWidget
	shaderProgram uint32
}

const (
	vertexShaderSource = `
    #version 460

    in vec3 vp;
	in vec2 vertexUV;

	out vec2 UV;

    void main() {
        gl_Position = vec4(vp, 1.0);

		UV = vertexUV;
    }
` + "\x00"

	fragmentShaderSource = `
	#version 460

	in vec2 UV;

	uniform sampler2D tex;

	out vec4 outputColor;

	void main() {
		outputColor = texture(tex, UV);
	}
` + "\x00"
)

func (g *Game) SetupGamePanel() {
	resizedQuad := make([]float32, len(quad))

	copy(resizedQuad, quad)

	panelWidth := float32(2)
	windowHeight := panelWidth

	for i := 0; i < len(resizedQuad); i += 3 {
		x := resizedQuad[i]
		y := resizedQuad[i+1]

		resizedQuad[i] = -1 + x*panelWidth
		resizedQuad[i+1] = (1 - windowHeight) + y*windowHeight
	}

	g.quadVao, g.UVBO = makeVao(resizedQuad, quadUVs)
}

func (g *Game) Render() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	draw.SwitchProgram(g.shaderProgram)
	gl.BindVertexArray(g.quadVao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, g.textureHandle)
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
	draw.RestoreGLOptions()
}

func (g *Game) InitGL() {
	g.shaderProgram = shader.CompileProgram(vertexShaderSource, fragmentShaderSource)

	gl.GenTextures(1, &g.textureHandle)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, g.textureHandle)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, 0)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, 0)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.NEAREST)

	g.SetupGamePanel()
}

func (g Game) OnWindowResize(w *Window) {
	g.SetupGamePanel()
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
