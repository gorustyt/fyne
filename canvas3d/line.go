package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Lines)(nil)
var _ gl.Canvas3DPainter = (*Lines)(nil)

type Line struct {
	Pos      *mgl32.Vec3
	Colorubs []uint8
}

type Lines struct {
	width float32
	lines []*Line
}

func NewLine() *Lines {
	return &Lines{width: 1}
}

func (l *Lines) AppendLine(v *Line) {
	l.lines = append(l.lines, v)
}

func (l *Lines) SetLineWidth(w float32) *Lines {
	l.width = w
	return l
}

func (l *Lines) Draw(p *gl.Painter3D, pos fyne.Position, Frame fyne.Size) {
	p.ExtPointSize(l.width)
	p.ExtBegin(gl.Lines)
	for _, v := range l.lines {
		p.ExtColor4ubv(v.Colorubs)
		p.ExtVertex3fV(*v.Pos)
	}
	p.ExtEnd()
}

func (l *Lines) InitOnce(p *gl.Painter3D) {

}

func (l *Lines) Init(p *gl.Painter3D) {

}

func (l *Lines) After(p *gl.Painter3D) {

}

func (l *Lines) NeedShader() bool {
	return false
}
