package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"strings"
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
}

var quad = []float32{
	-1, 1, -1,
	-1, -1, -1,
	1, -1, -1,

	-1, 1, -1,
	1, -1, -1,
	1, 1, -1,
}

const (
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"
)

func (w *Window) Open(resX, resY int) {
	w.Initialised = make(chan bool)

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

	w.glfwWindow, err = glfw.CreateWindow(resX, resY, "Testing", nil, nil)

	w.glfwWindow.SetAspectRatio(resX, resY)

	if err != nil {
		panic(err)
	}

	w.glfwWindow.MakeContextCurrent()

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

	go func() {
		w.Initialised <- true
	}()

	//gl.GenTextures(1, &w.Game.textureHandle)

	w.Game.quadVao = makeVao(quad)

	w.Game.Pixels = make([]uint8, resX*resY*3)

	i := 0

	for !w.glfwWindow.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(w.glProgramHandle)

		if i >= 3 {
			w.Game.Pixels[i-1] = 0x00
			w.Game.Pixels[i-2] = 0x00
			w.Game.Pixels[i-3] = 0x00
		}

		w.Game.Pixels[i] = 0xFF
		w.Game.Pixels[i+1] = 0xFF
		w.Game.Pixels[i+2] = 0xFF
		//Clear screen
		//for(int y = 0; y < SCREEN_HEIGHT; ++y)
		//for(int x = 0; x < SCREEN_WIDTH; ++x)
		//screenData[y][x][0] = screenData[y][x][1] = screenData[y][x][2] = 0;
		//

		//Create a texture
		//glTexImage2D(GL_TEXTURE_2D, 0, 3, SCREEN_WIDTH, SCREEN_HEIGHT, 0, GL_RGB, GL_UNSIGNED_BYTE, (GLvoid*)screenData);
		//
		//Set up the texture
		//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);
		//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
		//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP);
		//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP);
		//
		//Enable textures
		//glEnable(GL_TEXTURE_2D);

		gl.BindVertexArray(w.Game.quadVao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quad)/3))

		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
		i += 3
		i %= len(w.Game.Pixels)
	}
}

func (w *Window) Close() {
	w.glfwWindow.SetShouldClose(true)
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
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
	return &Window{}
}
