package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Fog)(nil)
var _ gl.Canvas3DPainter = (*Fog)(nil)

type Fog struct {
	Start   float32
	End     float32
	Colorf  mgl32.Vec4
	enable  uint8
	mode    int32
	Density float32
	hint    uint32
}

func NewFog() *Fog {
	return &Fog{enable: 1, mode: gl.Linear, hint: gl.FogDontCare}
}
func (f *Fog) Enable() {
	f.enable = 1
}

func (f *Fog) Disable() {
	if f.enable == 1 {
		f.enable = 2
	} else {
		f.enable = 0
	}
}
func (f *Fog) InitOnce(p *gl.Painter3D) {

}

func (f *Fog) Init(p *gl.Painter3D) {

}

func (f *Fog) After(p *gl.Painter3D) {

}

func (f *Fog) NeedShader() bool {
	return false
}
func (f *Fog) SetModeExp() {
	f.mode = gl.Exp
}

func (f *Fog) SetModeLinear() {
	f.mode = gl.Linear
}
func (f *Fog) Draw(p *gl.Painter3D, pos fyne.Position, Frame fyne.Size) {
	switch f.enable {
	case 0:

	case 1:
		p.Enable(gl.Fog)
		p.ExtFogi(gl.FogMode, f.mode)
		p.ExtFogf(gl.FogDensity, f.Density)
		p.ExtHint(gl.FogHint, f.hint)
		p.ExtFogf(gl.FogStart, f.Start)
		p.ExtFogf(gl.FogEnd, f.End)
		p.ExtFogfv(gl.FogColor, f.Colorf)
	case 2:
		f.enable = 0
		p.Disable(gl.Fog)
	}
}
