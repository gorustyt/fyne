package canvas3d

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Lines)(nil)
var _ gl.Canvas3DPainter = (*Lines)(nil)

type Line struct {
	P1 Point
	P2 Point
}

type Lines struct {
	width    float32
	lines    []*Line
	drawMode uint32
}

func NewLines() *Lines {
	return &Lines{width: 1, drawMode: gl.Lines}
}

func (l *Lines) AppendLine(v *Line) {
	l.lines = append(l.lines, v)
}

func (l *Lines) SetLineWidth(w float32) *Lines {
	l.width = w
	return l
}

func (l *Lines) SetDrawStrIp() {
	l.drawMode = gl.LineStrIp
}

func (l *Lines) SetDrawLoop() {
	l.drawMode = gl.LineLoop
}
func (l *Lines) Draw(p *gl.Painter3D, pos fyne.Position, Frame fyne.Size) {
	p.ExtLineWidth(l.width)
	p.ExtBegin(l.drawMode)
	for _, v := range l.lines {
		p.ExtColor4ubv(v.P1.Colorubs)
		if v.P1.Pos != nil {
			p.ExtVertex3fV(*v.P1.Pos)
		}
		p.ExtColor4ubv(v.P2.Colorubs)
		if v.P2.Pos != nil {
			p.ExtVertex3fV(*v.P2.Pos)
		}

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
