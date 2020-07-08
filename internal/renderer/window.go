package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	_ "image/png"
	"runtime"
	"strings"
	"unsafe"
)

type Window struct {
	glfwWindow      *glfw.Window
	glProgramHandle uint32
	Initialised     chan bool

	Game Game
}

type Game struct {
	Pixels        []uint8
	textureHandle uint32
	quadVao       uint32
	UVBO          uint32
	pixelsPointer unsafe.Pointer
}

var quad = []float32{
	-1, 1, 0, // Top Left
	-1, -1, 0, // Bottom Left
	1, -1, 0, // Bottom Right

	-1, 1, 0, // Top Left
	1, -1, 0, // Bottom Right
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

//	0, 1, // Top Left
//	0, 0, // Bottom Left
//	1, 0, // Bottom Right
//
//	0, 1, // Top Left
//	1, 0, // Bottom Right
//	1, 1, // Top Right

const (
	vertexShaderSource = `
    #version 410

	layout(location = 1) in vec2 vertexUV;

    in vec3 vp;

	out vec2 UV;

    void main() {
        gl_Position = vec4(vp, 1.0);

		UV = vertexUV;
    }
` + "\x00"

	fragmentShaderSource = `
	#version 410

	in vec2 UV;

	uniform sampler2D tex;

	out vec4 outputColor;

	void main() {
		outputColor = texture(tex, UV);
		//outputColor = vec4(UV.x, UV.y,0, 1);
		//outputColor = vec4(1,1,1,1);
	}
` + "\x00"
)

func (w *Window) Open(resX, resY, scale int, pixelPointer unsafe.Pointer) {
	runtime.LockOSThread()

	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	defer glfw.Terminate()

	w.glfwWindow, err = glfw.CreateWindow(resX * scale, resY * scale, "Jamboy", nil, nil)

	w.glfwWindow.SetAspectRatio(resX, resY)

	if err != nil {
		panic(err)
	}

	w.glfwWindow.MakeContextCurrent()
	w.glfwWindow.SetSizeCallback(w.onWindowSetSize)

	if err := gl.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialise opengl"))
	}

	w.glProgramHandle = gl.CreateProgram()

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)

	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	if err != nil {
		panic(err)
	}

	gl.AttachShader(w.glProgramHandle, vertexShader)
	gl.AttachShader(w.glProgramHandle, fragmentShader)
	gl.LinkProgram(w.glProgramHandle)

	w.Game.quadVao, w.Game.UVBO = makeVao(quad, quadUVs)

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

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	go func() {
		w.Initialised <- true
	}()

	i := 0

	for !w.glfwWindow.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		//for j := 0; j < resX*4; j += 4 {
		//	pi := (i * resX * 4) + j
		//	w.Game.Pixels[pi] = 0xFF
		//	w.Game.Pixels[pi+1] = 0x10
		//	w.Game.Pixels[pi+2] = 0x08
		//	w.Game.Pixels[pi+3] ^= 0xFF
		//}

		//copy(buffer, w.Game.Pixels)

		gl.UseProgram(w.glProgramHandle)
		gl.BindVertexArray(w.Game.quadVao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, w.Game.textureHandle)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(resX),
			int32(resY),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			pixelPointer,
		)

		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quad)/3))

		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()

		i++
		i %= resY
	}
}

func (w *Window) Close() {
	w.glfwWindow.SetShouldClose(true)
}

func (w *Window) onWindowSetSize(glfwWindow *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
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

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func NewWindow() *Window {
	window := &Window{
		Initialised: make(chan bool),
	}

	return window
}
