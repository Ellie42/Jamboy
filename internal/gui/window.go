package gui

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	gooey "git.agehadev.com/elliebelly/gooey/pkg"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	_ "image/png"
	"runtime"
	"unsafe"
)

type Window struct {
	glfwWindow      *glfw.Window
	glProgramHandle uint32
	Initialised     chan bool

	Game Game
	GUI  *gooey.Gooey
}

var (
	ColourBGGrey          = draw.NewRGBAFromHex("37505C")
	ColourBGGreyLight     = draw.NewRGBAFromHex("49626E")
	ColourBGGreyDark      = draw.NewRGBAFromHex("1E3440")
	ColourBGBlueDark      = draw.NewRGBAFromHex("16161a")
	ColourBGBlueVeryDark  = draw.NewRGBAFromHex("0e0e10")
	ColourBGBlueVeryDark2 = draw.NewRGBAFromHex("09090a")
	ColourBGPink          = draw.RGBA{0.741, 0.647, 0.608, 1}
	ColourBGPinkDark      = draw.RGBA{0.741, 0.576, 0.51, 1}
	ColourPaleRed         = draw.NewRGBAFromHex("A14643")
	ColourPalerRed        = draw.NewRGBAFromHex("B25D5A")
	ColourBrightRed       = draw.NewRGBAFromHex("962120")
	ColourDarkRed         = draw.NewRGBAFromHex("6e1817")
	ColourHighlight       = draw.RGBA{0.655, 0.322, 0.012, 1}
)

func (w *Window) Open(resX, resY int, pixelPointer unsafe.Pointer, pixels []uint8, jamboy *engine.Jamboy) {
	runtime.LockOSThread()

	w.Game.Pixels = pixels
	w.Game.gameBoyAspectRatioMulti = float32(resY) / float32(resX)
	w.Game.ResX = int32(resX)
	w.Game.ResY = int32(resY)
	w.Game.pixelsPointer = pixelPointer

	gui, err := gooey.Init()

	if err != nil {
		panic(err)
	}

	w.GUI = gui

	gooeyWindow, err := w.GUI.Window.CreateWindow(widget.WindowPreferences{
		Width:  1200,
		Height: 800,
		Title:  "Jamboy - Getting Gooey",
		GLFWHints: map[glfw.Hint]int{
			glfw.Resizable:               glfw.True,
			glfw.OpenGLProfile:           glfw.OpenGLCoreProfile,
			glfw.OpenGLForwardCompatible: glfw.True,
		},
		OpenedCB: func() {
			w.Initialised <- true
		},
	})

	if err != nil {
		panic(err)
	}

	//gooeyWindow.Context.ShowDebug = true

	gooeyWindow.Layout = widget.NewFreeLayout(
		&settings.WidgetPreferences{
			StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourBGBlueVeryDark},
		},
	)

	gamePanel := widget.NewPanel(&settings.WidgetPreferences{
		FixedRatioAxis:      settings.FixedY,
		FixedRatio:          0.9,
		AlignmentHorizontal: settings.HorizontalLeft,
		AlignmentVertical:   settings.VerticalTop,
		StyleSettings: &styles.StyleSettings{
			BackgroundColour: &ColourBrightRed,
		},
		Padding: &dimension.DirectionalRectSized{
			dimension.Size{0, dimension.SizeUnitPixels},
			dimension.Size{0, dimension.SizeUnitPixels},
			dimension.Size{0, dimension.SizeUnitPixels},
			dimension.Size{0, dimension.SizeUnitPixels},
		},
		DimensionBounds: &dimension.Dimensions{
			Width: &dimension.Size{
				400,
				dimension.SizeUnitPixels,
			},
		},
	},
		NewGameWidget(&w.Game),
	)

	debugger := NewDebuggerControl(jamboy, &gooeyWindow.Context.EventManager)

	codeListWidget := NewCodeListWidget(jamboy, debugger)

	gooeyWindow.Layout.AddChild(
		widget.NewLinearLayout(&settings.WidgetPreferences{
			Padding: &dimension.DirectionalRectSized{
				dimension.Size{Amount: 8},
				dimension.Size{Amount: 8},
				dimension.Size{Amount: 8},
				dimension.Size{Amount: 8},
			},
		},
			gamePanel,
			widget.NewPanel(nil,
				codeListWidget,
			),
		),
	)

	w.GUI.Loop()
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32, uvs []float32) (vao, uvbo uint32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	// Vertex data
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// UVData
	gl.GenBuffers(1, &uvbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(uvs), gl.Ptr(uvs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	return
}

func NewWindow() *Window {
	w := &Window{
		Initialised: make(chan bool),
	}

	return w
}
