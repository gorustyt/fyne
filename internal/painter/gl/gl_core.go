//go:build (!gles && !arm && !arm64 && !android && !ios && !mobile && !js && !test_web_driver && !wasm) || (darwin && !mobile && !ios && !js && !wasm && !test_web_driver)
// +build !gles,!arm,!arm64,!android,!ios,!mobile,!js,!test_web_driver,!wasm darwin,!mobile,!ios,!js,!wasm,!test_web_driver

package gl

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"image"
	"image/draw"
	"strings"

	"github.com/go-gl/gl/v4.2-compatibility/gl"

	"github.com/gorustyt/fyne/v2"
)

const (
	arrayBuffer           = gl.ARRAY_BUFFER
	bitColorBuffer        = gl.COLOR_BUFFER_BIT
	bitDepthBuffer        = gl.DEPTH_BUFFER_BIT
	clampToEdge           = gl.CLAMP_TO_EDGE
	colorFormatRGBA       = gl.RGBA
	compileStatus         = gl.COMPILE_STATUS
	constantAlpha         = gl.CONSTANT_ALPHA
	float                 = gl.FLOAT
	fragmentShader        = gl.FRAGMENT_SHADER
	front                 = gl.FRONT
	GlFalse               = gl.FALSE
	GlTrue                = gl.TRUE
	linkStatus            = gl.LINK_STATUS
	one                   = gl.ONE
	oneMinusConstantAlpha = gl.ONE_MINUS_CONSTANT_ALPHA
	oneMinusSrcAlpha      = gl.ONE_MINUS_SRC_ALPHA
	scissorTest           = gl.SCISSOR_TEST
	srcAlpha              = gl.SRC_ALPHA
	staticDraw            = gl.STATIC_DRAW
	texture0              = gl.TEXTURE0
	Texture2D             = gl.TEXTURE_2D
	TextureMinFilter      = gl.TEXTURE_MIN_FILTER
	TextureMagFilter      = gl.TEXTURE_MAG_FILTER
	textureWrapS          = gl.TEXTURE_WRAP_S
	textureWrapT          = gl.TEXTURE_WRAP_T
	Triangles             = gl.TRIANGLES
	triangleStrip         = gl.TRIANGLE_STRIP
	unsignedByte          = gl.UNSIGNED_BYTE
	vertexShader          = gl.VERTEX_SHADER
	GlRgba                = gl.RGBA
	GlUnsigedBytes        = gl.UNSIGNED_BYTE
	LinearMipMapNearest   = gl.LINEAR_MIPMAP_NEAREST
)

const noBuffer = context.Buffer(0)
const noShader = context.Shader(0)

var textureFilterToGL = []int32{gl.LINEAR, gl.NEAREST, gl.LINEAR}

func (p *painter) Init() {
	p.ctx = &coreContext{}
	err := gl.Init()
	if err != nil {
		fyne.LogError("failed to initialise OpenGL", err)
		return
	}

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	p.logError()
	p.program = p.createProgram("simple")
	p.lineProgram = p.createProgram("line")
	p.rectangleProgram = p.createProgram("rectangle")
	p.roundRectangleProgram = p.createProgram("round_rectangle")
}

type coreContext struct{}

var _ context.Context = (*coreContext)(nil)

func (c *coreContext) ActiveTexture(textureUnit uint32) {
	gl.ActiveTexture(textureUnit)
}

func (c *coreContext) AttachShader(program context.Program, shader context.Shader) {
	gl.AttachShader(uint32(program), uint32(shader))
}

func (c *coreContext) BindBuffer(target uint32, buf context.Buffer) {
	gl.BindBuffer(target, uint32(buf))
}

func (c *coreContext) BindTexture(target uint32, texture context.Texture) {
	gl.BindTexture(target, uint32(texture))
}

func (c *coreContext) BlendColor(r, g, b, a float32) {
	gl.BlendColor(r, g, b, a)
}

func (c *coreContext) BlendFunc(srcFactor, destFactor uint32) {
	gl.BlendFunc(srcFactor, destFactor)
}

func (c *coreContext) BufferData(target uint32, points []float32, usage uint32) {
	gl.BufferData(target, 4*len(points), gl.Ptr(points), usage)
}

func (c *coreContext) Clear(mask uint32) {
	gl.Clear(mask)
}

func (c *coreContext) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (c *coreContext) CompileShader(shader context.Shader) {
	gl.CompileShader(uint32(shader))
}

func (c *coreContext) CreateBuffer() context.Buffer {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	return context.Buffer(vbo)
}

func (c *coreContext) CreateProgram() context.Program {
	return context.Program(gl.CreateProgram())
}

func (c *coreContext) CreateShader(typ uint32) context.Shader {
	return context.Shader(gl.CreateShader(typ))
}

func (c *coreContext) CreateTexture() (texture context.Texture) {
	var tex uint32
	gl.GenTextures(1, &tex)
	return context.Texture(tex)
}

func (c *coreContext) DeleteBuffer(buffer context.Buffer) {
	gl.DeleteBuffers(1, (*uint32)(&buffer))
}

func (c *coreContext) DeleteTexture(texture context.Texture) {
	tex := uint32(texture)
	gl.DeleteTextures(1, &tex)
}

func (c *coreContext) Disable(capability uint32) {
	gl.Disable(capability)
}

func (c *coreContext) DrawElementsArrays(mode uint32, index []uint32) {
	gl.DrawElements(mode, int32(len(index)), gl.UNSIGNED_INT, gl.Ptr(index))
}

func (c *coreContext) DrawArrays(mode uint32, first, count int) {
	gl.DrawArrays(mode, int32(first), int32(count))
}

func (c *coreContext) Enable(capability uint32) {
	gl.Enable(capability)
}

func (c *coreContext) EnableVertexAttribArray(attribute context.Attribute) {
	gl.EnableVertexAttribArray(uint32(attribute))
}

func (c *coreContext) GetAttribLocation(program context.Program, name string) context.Attribute {
	return context.Attribute(gl.GetAttribLocation(uint32(program), gl.Str(name+"\x00")))
}

func (c *coreContext) GetError() uint32 {
	return gl.GetError()
}

func (c *coreContext) GetProgrami(program context.Program, param uint32) int {
	var value int32
	gl.GetProgramiv(uint32(program), param, &value)
	return int(value)
}

func (c *coreContext) GetProgramInfoLog(program context.Program) string {
	var logLength int32
	gl.GetProgramiv(uint32(program), gl.INFO_LOG_LENGTH, &logLength)
	info := strings.Repeat("\x00", int(logLength+1))
	gl.GetProgramInfoLog(uint32(program), logLength, nil, gl.Str(info))
	return info
}

func (c *coreContext) GetShaderi(shader context.Shader, param uint32) int {
	var value int32
	gl.GetShaderiv(uint32(shader), param, &value)
	return int(value)
}

func (c *coreContext) GetShaderInfoLog(shader context.Shader) string {
	var logLength int32
	gl.GetShaderiv(uint32(shader), gl.INFO_LOG_LENGTH, &logLength)
	info := strings.Repeat("\x00", int(logLength+1))
	gl.GetShaderInfoLog(uint32(shader), logLength, nil, gl.Str(info))
	return info
}

func (c *coreContext) GetUniformLocation(program context.Program, name string) context.Uniform {
	return context.Uniform(gl.GetUniformLocation(uint32(program), gl.Str(name+"\x00")))
}

func (c *coreContext) LinkProgram(program context.Program) {
	gl.LinkProgram(uint32(program))
}

func (c *coreContext) ReadBuffer(src uint32) {
	gl.ReadBuffer(src)
}

func (c *coreContext) ReadPixels(x, y, width, height int, colorFormat, typ uint32, pixels []uint8) {
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), colorFormat, typ, gl.Ptr(pixels))
}

func (c *coreContext) Scissor(x, y, w, h int32) {
	gl.Scissor(x, y, w, h)
}

func (c *coreContext) ShaderSource(shader context.Shader, source string) {
	csources, free := gl.Strs(source + "\x00")
	defer free()
	gl.ShaderSource(uint32(shader), 1, csources, nil)
}

func (c *coreContext) TexImage2D(target uint32, level, width, height int, colorFormat, typ uint32, data []uint8) {
	gl.TexImage2D(
		target,
		int32(level),
		int32(colorFormat),
		int32(width),
		int32(height),
		0,
		colorFormat,
		typ,
		gl.Ptr(data),
	)
}

func (c *coreContext) TexParameteri(target, param uint32, value int32) {
	gl.TexParameteri(target, param, value)
}

func (c *coreContext) Uniform1f(uniform context.Uniform, v float32) {
	gl.Uniform1f(int32(uniform), v)
}

func (c *coreContext) Uniform2f(uniform context.Uniform, v0, v1 float32) {
	gl.Uniform2f(int32(uniform), v0, v1)
}
func (c *coreContext) Uniform3f(uniform context.Uniform, v mgl32.Vec3) {
	gl.Uniform3f(int32(uniform), v[0], v[1], v[2])
}

func (c *coreContext) Uniform4f(uniform context.Uniform, v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(uniform), v0, v1, v2, v3)
}

func (c *coreContext) UseProgram(program context.Program) {
	gl.UseProgram(uint32(program))
}

func (c *coreContext) VertexAttribPointerWithOffset(attribute context.Attribute, size int, typ uint32, normalized bool, stride, offset int) {
	gl.VertexAttribPointerWithOffset(uint32(attribute), int32(size), typ, normalized, int32(stride), uintptr(offset))
}

func (c *coreContext) Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (c *coreContext) UniformMatrix4fv(program context.Program, name string, mat4 mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(c.GetUniformLocation(program, name)), 1, false, &mat4[0])
}

func (c *coreContext) Uniform1i(program context.Program, name string, v0 int32) {
	gl.Uniform1i(int32(c.GetUniformLocation(program, name)), v0)
}

func (c *coreContext) EnableDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (c *coreContext) DisableDepthTest() {
	gl.Disable(gl.DEPTH_TEST)
}

func (c *coreContext) MakeVaoWithEbo(points []float32, indexs []uint32) (context.Buffer, context.Buffer) {
	vbo := c.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(vbo))
	c.BufferData(gl.ARRAY_BUFFER, points, gl.STATIC_DRAW)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indexs), gl.Ptr(indexs), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return vbo, context.Buffer(ebo)
}

func (c *coreContext) MakeVao(points []float32) context.Buffer {
	var vbo uint32

	// 在显卡中开辟一块空间，创建顶点缓存对象，个数为1，变量vbo会被赋予一个ID值。
	gl.GenBuffers(1, &vbo)

	// 将 vbo 赋值给 gl.ARRAY_BUFFER，要知道这个对象会被赋予不同的vbo，因此其值是变化的
	// 可选类型：GL_ARRAY_BUFFER, GL_ELEMENT_ARRAY_BUFFER, GL_PIXEL_PACK_BUFFER, GL_PIXEL_UNPACK_BUFFER
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// 将内存中的数据传递到显卡中的gl.ARRAY_BUFFER对象上，其实是把数据传递到绑定在其上面的vbo对象上。
	// 4*len(points) 代表总的字节数，因为是32位的
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	// 创建顶点数组对象，个数为1，变量vao会被赋予一个ID值。
	gl.GenVertexArrays(1, &vao)
	// 后面的两个函数都是要操作具体的vao的，因此需要先将vao绑定到opengl上。
	// 解绑：gl.BindVertexArray(0)，opengl中很多的解绑操作都是传入0
	gl.BindVertexArray(vao)
	return context.Buffer(vbo)
}
func (c *coreContext) MakeTexture(img image.Image, index uint32) context.Texture {
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	var te uint32
	gl.GenTextures(1, &te)
	gl.ActiveTexture(index)
	gl.BindTexture(gl.TEXTURE_2D, te)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	return context.Texture(te)
}

func (c *coreContext) MakeLevelTexture(index uint32, size int32, levelData map[int32][]uint32) context.Texture {
	var te uint32
	gl.GenTextures(1, &te)
	gl.ActiveTexture(index)
	gl.BindTexture(gl.TEXTURE_2D, te)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	for k, v := range levelData {
		gl.TexImage2D(gl.TEXTURE_2D, k, gl.RGBA, size, size, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(v))
	}
	return context.Texture(te)
}

func GetTextureByIndex(index int) uint32 {
	switch index {
	case 0:
		return gl.TEXTURE0
	case 1:
		return gl.TEXTURE1
	case 2:
		return gl.TEXTURE2
	case 3:
		return gl.TEXTURE3
	case 4:
		return gl.TEXTURE4
	case 5:
		return gl.TEXTURE5
	case 6:
		return gl.TEXTURE6
	case 7:
		return gl.TEXTURE7
	case 8:
		return gl.TEXTURE8
	case 9:
		return gl.TEXTURE9
	case 10:
		return gl.TEXTURE10
	}
	return 0
}

func (c *coreContext) ExtBegin(mode uint32) {
	gl.Begin(mode)
}

func (c *coreContext) ExtEnd() {
	gl.End()
}

func (c *coreContext) ExtColor4ub(r, g, b, a uint8) {
	gl.Color4ub(r, g, b, a)
}

func (c *coreContext) ExtColor4ubv(v []uint8) {
	gl.Color4ubv(&v[0])
}

func (c *coreContext) ExtColor3fV(vec3 mgl32.Vec3) {
	gl.Color3fv(&vec3[0])
}
func (c *coreContext) ExtColor3f(r, g, b float32) {
	gl.Color3f(r, g, b)
}

func (c *coreContext) ExtPointSize(size float32) {
	gl.PointSize(size)
}

func (c *coreContext) ExtLineWidth(width float32) {
	gl.LineWidth(width)
}
func (c *coreContext) ExtFlush() {
	gl.Flush()
}

func (c *coreContext) ExtVertex3f(x, y, z float32) {
	gl.Vertex3f(x, y, z)
}

func (c *coreContext) ExtVertex3fV(vec3 mgl32.Vec3) {
	gl.Vertex3fv(&vec3[0])
}

const (
	LineStrIp   = gl.LINE_STRIP
	LineLoop    = gl.LINE_LOOP
	Quads       = gl.QUADS
	Lines       = gl.LINES
	Points      = gl.POINTS
	Fog         = gl.FOG
	FogMode     = gl.FOG_MODE
	FogDensity  = gl.FOG_DENSITY
	FogHint     = gl.FOG_HINT
	FogDontCare = gl.DONT_CARE
	Exp         = gl.EXP
	Exp2        = gl.EXP2
	FogStart    = gl.FOG_START
	FogEnd      = gl.FOG_END
	FogColor    = gl.FOG_COLOR
	CullFace    = gl.CULL_FACE
	Lequal      = gl.LEQUAL
	Linear      = gl.LINEAR
)

func (c *coreContext) ExtFogi(panme uint32, v int32) {
	gl.Fogi(panme, v)
}

func (c *coreContext) ExtFogiv(panme uint32, v []int32) {
	gl.Fogiv(panme, &v[0])
}

func (c *coreContext) GenTextures(n int32, textures *uint32) {
	gl.GenTextures(n, textures)
}
func (c *coreContext) ExtFogf(panme uint32, v float32) {
	gl.Fogf(panme, v)
}

func (c *coreContext) ExtFogfv(panme uint32, v mgl32.Vec4) {
	gl.Fogfv(panme, &v[0])
}

func (c *coreContext) ExtHint(target uint32, mode uint32) {
	gl.Hint(target, mode)
}

func (c *coreContext) ExtDepthMask(flag bool) {
	gl.DepthMask(flag)
}

func (c *coreContext) ExtVertex3fv(vec mgl32.Vec3) {
	gl.Vertex3fv(&vec[0])
}

func (c *coreContext) ExtTexCoord2fv(vec mgl32.Vec2) {
	gl.TexCoord2fv(&vec[0])
}

func (c *coreContext) ExtVertex2f(x, y float32) {
	gl.Vertex2f(x, y)
}

//func (c *coreContext) ExtGetIntegerv(pname uint32, data []int32) {
//	gl.GetIntegerv(pname, &data[0])
//}
//
//func (c *coreContext) ExtMatrixMode(mode uint32) {
//	gl.MatrixMode(mode)
//}
//
//func (c *coreContext) ExtRotatef(angle float32, x float32, y float32, z float32) {
//	gl.Rotatef(angle, x, y, z)
//}
//func (c *coreContext) ExtTranslatef(x float32, y float32, z float32) {
//	gl.Translatef(x, y, z)
//}
//func (c *coreContext) ExtGetDoublev(pname uint32, data []float64) {
//	gl.GetDoublev(pname, &data[0])
//}
//
//func (c *coreContext) ExtProject(x, y, z float64, model, project mgl64.Mat4, initialX, initialY, width, height int) mgl64.Vec3 {
//	return mgl64.Project(mgl64.Vec3{x, y, z}, model, project, initialX, initialY, width, height)
//}
//func (c *coreContext) ExtUnProject(x, y, z float64, model, project mgl64.Mat4, initialX, initialY, width, height int) (mgl64.Vec3, error) {
//	return mgl64.UnProject(mgl64.Vec3{x, y, z}, model, project, initialX, initialY, width, height)
//}
//
//func (c *coreContext) ExtLoadIdentity() {
//	gl.LoadIdentity()
//}
//
//func (c *coreContext) ExtOrtho2D(left, right, bottom, top float64) {
//	gl.Ortho(left, right, bottom, top, -1, 1)
//}
