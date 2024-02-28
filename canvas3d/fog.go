package canvas3d

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Fog)(nil)
var _ gl.Canvas3DPainter = (*Fog)(nil)

type Fog struct {
}

func (f Fog) InitOnce(p *gl.Painter3D) {

}

func (f Fog) Init(p *gl.Painter3D) {

}

func (f Fog) After(p *gl.Painter3D) {

}

func (f Fog) NeedShader() bool {
	return false
}

func (f Fog) Draw(p *gl.Painter3D, pos fyne.Position, Frame fyne.Size) {

}
