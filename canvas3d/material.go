package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Material)(nil)

type Material struct {
	Ambient   mgl32.Vec3
	Diffuse   mgl32.Vec3
	Specular  mgl32.Vec3
	Shininess float32

	AmbientStrength  float32
	DiffuseStrength  float32
	SpecularStrength float32
}

func (m *Material) InitOnce(p *gl.Painter3D) {

}

func (m *Material) Init(p *gl.Painter3D) {
	p.UniformVec3("material.ambient", m.Ambient)
	p.UniformVec3("material.diffuse", m.Diffuse)
	p.UniformVec3("material.specular", m.Specular)

	p.Uniform1f("material.ambient_strength", m.AmbientStrength)
	p.Uniform1f("material.diffuse_strength", m.DiffuseStrength)
	p.Uniform1f("material.specular_strength", m.SpecularStrength)
}

func (m *Material) After(p *gl.Painter3D) {

}

func NewMaterial() *Material {
	return &Material{
		Shininess:        1,
		AmbientStrength:  1,
		DiffuseStrength:  1,
		SpecularStrength: 1,
	}
}
