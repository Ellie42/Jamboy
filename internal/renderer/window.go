package renderer

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	gooey "git.agehadev.com/elliebelly/gooey/pkg"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"git.agehadev.com/elliebelly/gooey/pkg/widget/settings"
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

func (w *Window) Open(resX, resY int, pixelPointer unsafe.Pointer, pixels []uint8) {
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

	gooeyWindow, err := w.GUI.Window.CreateWindow(widget.Preferences{
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

	gooeyWindow.Context.ShowDebug = true

	panel := widget.NewPanel(&settings.WidgetPreferences{
		FixedRatioAxis: settings.FixedY,
		FixedRatio:     0.9,
		Padding: &dimension.DirectionalRect{
			Top:    0.01,
			Right:  0.01,
			Bottom: 0.01,
			Left:   0.01,
		},
		AlignmentHorizontal: settings.HorizontalMiddle,
		AlignmentVertical:   settings.VerticalMiddle,
	},
		NewGameWidget(&w.Game),
	)

	gooeyWindow.Layout.AddChild(
		widget.NewLinearLayout(&settings.WidgetPreferences{},
			panel,
			widget.NewPanel(nil),
		),
	)

	w.GUI.Loop()
}


func (w *Window) Close() {
	w.glfwWindow.SetShouldClose(true)
}

func (w *Window) onWindowSetSize(glfwWindow *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	w.Game.OnWindowResize(w)
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
