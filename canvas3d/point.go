package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Points)(nil)
var _ gl.Canvas3DPainter = (*Points)(nil)

type Point struct {
	Pos      *mgl32.Vec3
	Color    *mgl32.Vec3
	Colorubs []uint8
}

type Points struct {
	width float32
	ps    []*Point
}

func NewPoint() *Point {
	return &Point{
		Pos:   &mgl32.Vec3{},
		Color: &mgl32.Vec3{},
	}
}

func NewPoints() *Points {
	return &Points{
		width: 2,
	}
}

func (p2 *Points) AppendPoint(v *Point) {
	p2.ps = append(p2.ps, v)
}

func (p2 *Points) SetPointWidth(w float32) *Points {
	p2.width = w
	return p2
}
func (p2 *Points) Draw(p *gl.Painter3D, pos fyne.Position, Frame fyne.Size) {
	p.ExtPointSize(p2.width)
	p.ExtBegin(gl.Points)
	for _, v := range p2.ps {
		if v.Color != nil {
			p.ExtColor3fV(*v.Color)
		} else if v.Colorubs != nil {
			p.ExtColor4ubv(v.Colorubs[:])
		}
		p.ExtVertex3fV(*v.Pos)
	}
	p.ExtEnd()
}

func (p2 *Points) InitOnce(p *gl.Painter3D) {

}

func (p2 *Points) Init(p *gl.Painter3D) {

}

func (p2 *Points) After(p *gl.Painter3D) {

}

func (p2 *Points) NeedShader() bool {
	return false
}
