package canvas3d

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

type ICanvas3d interface {
	SetShaderConfig(vertStr, fragStr string)
	AppendObj(obj gl.Canvas3D)
	Reset()
	GetRenderObj() fyne.CanvasObject
	AppendRenderFunc(fn func(painter context.Painter))
}
type Canvas3d struct {
	*gl.Canvas3dObjs
}

func NewCanvas3d(n int) *Canvas3d {
	return &Canvas3d{Canvas3dObjs: gl.NewCustomObjs(n)}
}

func (c *Canvas3d) AppendRenderFunc(index int, fn func(ctx context.Painter)) {
	obj := c.GetCanvas3dObj(index)
	obj.RenderFuncs = append(obj.RenderFuncs, fn)
}

func (c *Canvas3d) AppendObj(index int, obj gl.Canvas3D) {
	o := c.GetCanvas3dObj(index)
	o.Objs = append(o.Objs, obj)
}
func (c *Canvas3d) SetShaderConfig(index int, vertStr, fragStr string) {
	o := c.GetCanvas3dObj(index)
	o.ChangeShader(vertStr, fragStr)
}
func (c *Canvas3d) Reset() {
	c.RangeCanvas3dObj(func(obj *gl.Canvas3dObj) (stop bool) {
		obj.Objs = obj.Objs[:0]
		obj.RenderFuncs = obj.RenderFuncs[:0]
		return false
	})

}

func (c *Canvas3d) GetRenderObj() fyne.CanvasObject {
	return c.Canvas3dObjs
}
