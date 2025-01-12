package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

const (
	LightConstant    = "light.constant"
	LightLinear      = "light.linear"
	LightQuadratic   = "light.quadratic"
	LightCutOff      = "light.cutOff"
	LightOuterCutOff = "light.outerCutOff"
	LightDirection   = "light.direction"
	LightPosition    = "light.position"
)

type Light struct {
	Position  mgl32.Vec3
	Direction mgl32.Vec3
	Constant  float32
	Linear    float32
	Quadratic float32
	*Material
}

func (m *Light) Init(p *gl.Painter3D) {
	m.Material.Init(p)
	p.UniformVec3(LightPosition, m.Position)
	p.UniformVec3(LightDirection, m.Direction)
	p.Uniform1f(LightConstant, m.Constant)
	p.Uniform1f(LightLinear, m.Linear)
	p.Uniform1f(LightQuadratic, m.Quadratic)
}

func (m *Light) After(p *gl.Painter3D) {

}

func NewLight() *Light {
	l := &Light{
		Material: NewMaterial(),
	}
	l.Name = "light"
	return l
}

type PointLight struct {
	*Light
}

func NewPointLight() *PointLight {
	return &PointLight{
		Light: NewLight(),
	}
}
func (m *PointLight) Init(p *gl.Painter3D) {
	m.Light.Init(p)
}

func (m *PointLight) After(p *gl.Painter3D) {

}

type SpotLight struct {
	*Light
	CutOff      float32
	OuterCutOff float32
}

func NewSpotLight() *SpotLight {
	return &SpotLight{
		Light: NewLight(),
	}
}
func (m *SpotLight) Init(p *gl.Painter3D) {
	m.Light.Init(p)
	p.Uniform1f(LightCutOff, m.CutOff)
	p.Uniform1f(LightOuterCutOff, m.OuterCutOff)
}

func (m *SpotLight) After(p *gl.Painter3D) {

}

type DirectionLight struct {
	*Light
}

func (m *DirectionLight) Init(p *gl.Painter3D) {
	m.Light.Init(p)
}
func (m *DirectionLight) After(p *gl.Painter3D) {

}
func NewDirectionLight() *DirectionLight {
	return &DirectionLight{
		Light: NewLight(),
	}
}
