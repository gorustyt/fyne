package canvas3d

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Material)(nil)

const (
	materialAttr = "material"
)

type Material struct {
	Ambient          mgl32.Vec3
	Diffuse          mgl32.Vec3
	Specular         mgl32.Vec3
	Shininess        float32
	Name             string
	AmbientStrength  float32
	DiffuseStrength  float32
	SpecularStrength float32
}

func (m *Material) NeedShader() bool {
	return true
}

func (m *Material) InitOnce(p *gl.Painter3D) {

}
func (m *Material) GetName(attr string) string {
	return fmt.Sprintf("%v.%v", m.Name, attr)
}
func (m *Material) Init(p *gl.Painter3D) {
	p.UniformVec3(m.GetName("ambient"), m.Ambient)
	p.UniformVec3(m.GetName("diffuse"), m.Diffuse)
	p.UniformVec3(m.GetName("specular"), m.Specular)
	if m.Name == materialAttr {
		p.Uniform1f(m.GetName("shininess"), m.Shininess)
	}
	p.Uniform1f(m.GetName("ambient_strength"), m.AmbientStrength)
	p.Uniform1f(m.GetName("diffuse_strength"), m.DiffuseStrength)
	p.Uniform1f(m.GetName("specular_strength"), m.SpecularStrength)
}

func (m *Material) After(p *gl.Painter3D) {

}

func NewMaterial() *Material {
	return &Material{
		Name:             materialAttr,
		Shininess:        1,
		AmbientStrength:  1,
		DiffuseStrength:  1,
		SpecularStrength: 1,
	}
}
