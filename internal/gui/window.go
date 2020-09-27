package gui

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	gooey "git.agehadev.com/elliebelly/gooey/pkg"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/styles"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/go-gl/gl/v4.6-compatibility/gl"
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
	ColourGrey          = draw.NewRGBAFromHex("37505C")
	ColourGreyLight     = draw.NewRGBAFromHex("49626E")
	ColourGreyDark      = draw.NewRGBAFromHex("252528")
	ColourBlue          = draw.NewRGBAFromHex("314FDA")
	ColourBlueDark      = draw.NewRGBAFromHex("16161a")
	ColourBlueVeryDark  = draw.NewRGBAFromHex("0e0e10")
	ColourBlueVeryDark2 = draw.NewRGBAFromHex("09090a")
	ColourPink          = draw.RGBA{0.741, 0.647, 0.608, 1}
	ColourGreen         = draw.NewRGBAFromHex("37E13E")
	ColourPinkDark      = draw.RGBA{0.741, 0.576, 0.51, 1}
	ColourPaleRed       = draw.NewRGBAFromHex("A14643")
	ColourPalerRed      = draw.NewRGBAFromHex("B25D5A")
	ColourBrightRed     = draw.NewRGBAFromHex("962120")
	ColourDarkRed       = draw.NewRGBAFromHex("6e1817")
	ColourRed           = draw.NewRGBAFromHex("CD2524")
	ColourPurple        = draw.NewRGBAFromHex("6A2AD8")
	ColourHighlight     = draw.RGBA{0.655, 0.322, 0.012, 1}
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
			StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourBlueVeryDark},
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
				float32(w.Game.ResX) * 2,
				dimension.SizeUnitPixels,
			},
		},
	},
		NewGameWidget(&w.Game),
	)

	debugger := NewDebuggerControl(jamboy, &gooeyWindow.Context.EventManager)

	codeListWidget := NewCodeListWidget(jamboy, debugger)

	registerRowPrefs := &settings.WidgetPreferences{
		DimensionBounds: &dimension.Dimensions{
			Height: &dimension.Size{
				Amount: 32,
			},
		},
	}

	registerLayout := widget.NewLinearLayout(
		&settings.WidgetPreferences{
			Name:          "Registers",
			StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourGreyDark},
			DimensionBounds: &dimension.Dimensions{
				Height: &dimension.Size{
					Amount: 224,
				},
			},
		},
	)

	for i := 0; i < 14; i += 2 {
		var reg, reg2 = engine.RegisterID(i), engine.RegisterID(i + 1)

		registerLayout.AddChild(
			widget.NewLinearLayout(registerRowPrefs,
				widget.NewTextWidget(nil, engine.RegisterID(i).String()),
				widget.NewDynamicTextWidget(nil, func() string {
					return fmt.Sprintf("$%04x", jamboy.CPU.ReadRegisterInstant(reg))
				}),
				widget.NewTextWidget(nil, engine.RegisterID(i+1).String()),
				widget.NewDynamicTextWidget(nil, func() string {
					return fmt.Sprintf("$%04x", jamboy.CPU.ReadRegisterInstant(reg2))
				}),
			),
		)
	}

	registerLayout.Alignment = widget.LLAlignmentVertical

	rightSidePanel := widget.NewLinearLayout(
		&settings.WidgetPreferences{
			DimensionBounds: &dimension.Dimensions{
				Width: &dimension.Size{
					Amount: 300,
				},
			},
		},
		widget.NewPanel(nil),
		registerLayout,
	)
	rightSidePanel.Alignment = widget.LLAlignmentVertical

	//gooeyWindow.Layout.AddChild(getTestWindow())
	//w.GUI.Loop()
	//return

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
			rightSidePanel,
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

func getTestWindow() widget.Widget {
	//ll := widget.NewLinearLayout(nil,
	//	widget.NewPanel(&settings.WidgetPreferences{
	//		Name:          "PanelRed",
	//		StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourRed},
	//		DimensionBounds: &dimension.Dimensions{
	//			Width: &dimension.Size{
	//				Amount: 100,
	//			},
	//			Height: nil,
	//		},
	//	}),
	//	widget.NewPanel(&settings.WidgetPreferences{
	//		Name:          "PanelGreen",
	//		StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourGreen},
	//	}),
	//	widget.NewPanel(&settings.WidgetPreferences{
	//		Name:          "PanelBlue",
	//		StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourBlue},
	//		DimensionBounds: &dimension.Dimensions{
	//			Width: &dimension.Size{
	//				Amount: 100,
	//			},
	//			Height: nil,
	//		},
	//	}),
	//	widget.NewPanel(&settings.WidgetPreferences{
	//		Name:          "PanelPurple",
	//		StyleSettings: &styles.StyleSettings{BackgroundColour: &ColourPurple},
	//	}),
	//)

	list := widget.NewList(
		widget.NewListStringContentWidget,
		func(w widget.Widget, index int) {
			t := w.(*widget.Text)

			t.Text = fmt.Sprintf("Line %d is here", index)
		},
		nil,
	)

	return list
}
