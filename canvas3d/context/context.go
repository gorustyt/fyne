package context

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/cache"
	"image"
)

type Context interface {
	ActiveTexture(textureUnit uint32)
	AttachShader(program Program, shader Shader)
	BindBuffer(target uint32, buf Buffer)
	BindTexture(target uint32, texture Texture)
	BlendColor(r, g, b, a float32)
	BlendFunc(srcFactor, destFactor uint32)
	BufferData(target uint32, points []float32, usage uint32)
	Clear(mask uint32)
	ClearColor(r, g, b, a float32)
	CompileShader(shader Shader)
	CreateBuffer() Buffer
	CreateProgram() Program
	CreateShader(typ uint32) Shader
	CreateTexture() Texture
	DeleteBuffer(buffer Buffer)
	DeleteTexture(texture Texture)
	Disable(capability uint32)
	DrawArrays(mode uint32, first, count int)
	Enable(capability uint32)
	EnableVertexAttribArray(attribute Attribute)
	GetAttribLocation(program Program, name string) Attribute
	GetError() uint32
	GetProgrami(program Program, param uint32) int
	GetProgramInfoLog(program Program) string
	GetShaderi(shader Shader, param uint32) int
	GetShaderInfoLog(shader Shader) string
	GetUniformLocation(program Program, name string) Uniform
	LinkProgram(program Program)
	ReadBuffer(src uint32)
	ReadPixels(x, y, width, height int, colorFormat, typ uint32, pixels []uint8)
	Scissor(x, y, w, h int32)
	ShaderSource(shader Shader, source string)
	TexImage2D(target uint32, level, width, height int, colorFormat, typ uint32, data []uint8)
	TexParameteri(target, param uint32, value int32)
	Uniform1f(uniform Uniform, v float32)
	Uniform2f(uniform Uniform, v0, v1 float32)
	Uniform3f(uniform Uniform, vec3 mgl32.Vec3)
	Uniform4f(uniform Uniform, v0, v1, v2, v3 float32)
	UseProgram(program Program)
	VertexAttribPointerWithOffset(attribute Attribute, size int, typ uint32, normalized bool, stride, offset int)
	Viewport(x, y, width, height int)

	UniformMatrix4fv(program Program, name string, mat4 mgl32.Mat4)
	Uniform1i(program Program, name string, v0 int32)
	DisableDepthTest()
	EnableDepthTest()

	DrawElementsArrays(mode uint32, index []uint32)
	MakeTexture(img image.Image, index uint32) Texture
	MakeVao(points []float32) Buffer
	MakeVaoWithEbo(points []float32, index []uint32) (Buffer, Buffer)

	ExtFogi(panme uint32, v int32)
	ExtFogiv(panme uint32, v []int32)
	ExtFogf(panme uint32, v float32)
	ExtFogfv(panme uint32, v mgl32.Vec4)
	ExtBegin(mode uint32)
	ExtEnd()
	ExtColor3f(r, g, b float32)
	ExtColor3fV(vec3 mgl32.Vec3)
	ExtPointSize(size float32)
	ExtLineWidth(width float32)
	ExtFlush()
	ExtVertex3f(x, y, z float32)
	ExtVertex3fV(vec3 mgl32.Vec3)
	ExtColor4ub(r, g, b, a uint8)
	ExtColor4ubv(v []uint8)
}

type (
	// Texture represents an uploaded GL texture
	Texture cache.TextureType

	// Attribute represents a GL attribute
	Attribute int32
	// Buffer represents a GL buffer
	Buffer uint32
	// Program represents a compiled GL program
	Program uint32
	// Shader represents a GL shader
	Shader uint32
	// Uniform represents a GL uniform
	Uniform int32
)
