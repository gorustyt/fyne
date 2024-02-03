package canvas3d

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

type ICanvas3d interface {
	SetShaderConfig(vertStr, fragStr string)
	AppendObj(obj gl.Canvas3D)
	Reset()
	fyne.CanvasObject
}
type Canvas3d struct {
	*gl.Canvas3dObj
}

func NewCanvas3d() *Canvas3d {
	return &Canvas3d{}
}

func (c *Canvas3d) AppendObj(obj gl.Canvas3D) {
	c.Objs = append(c.Objs, obj)
}
func (c *Canvas3d) SetShaderConfig(vertStr, fragStr string) {
	c.VertStr, c.FragStr = vertStr, fragStr
}
func (c *Canvas3d) Reset() {
	c.Objs = c.Objs[:0]
}
