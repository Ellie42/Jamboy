package renderer

import (
	gooey "git.agehadev.com/elliebelly/gooey/pkg"
	"git.agehadev.com/elliebelly/gooey/pkg/window"
	"git.agehadev.com/elliebelly/gooey/pkg/window/widget"
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

	gooeyWindow, err := w.GUI.Window.CreateWindow(window.Preferences{
		Width:  2560,
		Height: 1440,
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

	gooeyWindow.Layout.Add(
		widget.NewPanel(
			NewGameWidget(&w.Game),
		),
	)

	w.GUI.Loop()

	//err := glfw.Init()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//w.glfwWindow, err = glfw.CreateWindow(2560, 1440, "Jamboy", nil, nil)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//w.glfwWindow.MakeContextCurrent()
	//w.glfwWindow.SetSizeCallback(w.onWindowSetSize)
	//
	//if err := gl.Init(); err != nil {
	//	panic(fmt.Sprintf("failed to initialise opengl"))
	//}
	//
	//w.glProgramHandle = gl.CreateProgram()
	//
	//vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//gl.AttachShader(w.glProgramHandle, vertexShader)
	//gl.AttachShader(w.glProgramHandle, fragmentShader)
	//gl.LinkProgram(w.glProgramHandle)
	//
	//w.Game.InitGL(w)
	//
	//gl.ClearColor(0.1, 0.1, 0.2, 1.0)
	//
	//go func() {
	//	w.Initialised <- true
	//}()
	//
	//for !w.glfwWindow.ShouldClose() {
	//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	//
	//	gl.UseProgram(w.glProgramHandle)
	//
	//	w.Game.Render(w)
	//
	//	w.glfwWindow.SwapBuffers()
	//	glfw.PollEvents()
	//}
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

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.GenBuffers(1, &uvbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(uvs), gl.Ptr(uvs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvbo)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	return
}

func NewWindow() *Window {
	w := &Window{
		Initialised: make(chan bool),
	}

	return w
}
