package context

import (
	"github.com/go-gl/mathgl/mgl32"
	"image"
)

type Painter interface {
	UniformMatrix4fv(name string, mat4 mgl32.Mat4)
	Uniform1i(name string, v0 int32)

	Uniform1f(name string, v float32)
	Uniform2f(name string, v0, v1 float32)
	Uniform3f(name string, v0, v1, v2 float32)
	UniformVec3(name string, vec3 mgl32.Vec3)
	Uniform4f(name string, v0, v1, v2, v3 float32)

	DrawElementsArrays(mode uint32, index []uint32)
	MakeTexture(img image.Image, index uint32) Texture
	MakeVao(points []float32) Buffer
	MakeVaoWithEbo(points []float32, index []uint32) (Buffer, Buffer)
}
