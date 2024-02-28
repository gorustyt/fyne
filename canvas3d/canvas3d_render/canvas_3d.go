package canvas3d_render

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/internal/driver/glfw"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

type ICanvas3d interface {
	SetShaderConfig(index int, vertStr, fragStr string)
	AppendObj(index int, obj gl.Canvas3D)
	AppendDefaultObj(obj gl.Canvas3D)
	Reset()
	GetRenderObj() fyne.CanvasObject
	AppendDefaultRenderFunc(index int, fn func(ctx context.Painter))
	AppendRenderFunc(index int, fn func(painter context.Painter))
	fyne.CanvasObject
}
type Canvas3d struct {
	*gl.Canvas3dObjs
}

func NewCanvas3d(n ...int) ICanvas3d {
	num := 1
	if len(n) > 0 {
		num = n[0]
	}
	if num <= 0 {
		num = 1
	}
	return &Canvas3d{Canvas3dObjs: gl.NewCustomObjs(num)}
}

func GetGlfwTime() float64 {
	return glfw.GetGlfwTime()
}

func (c *Canvas3d) AppendRenderFunc(index int, fn func(ctx context.Painter)) {
	obj := c.GetCanvas3dObj(index)
	obj.RenderFuncs = append(obj.RenderFuncs, fn)
}

func (c *Canvas3d) AppendDefaultRenderFunc(index int, fn func(ctx context.Painter)) {
	c.AppendRenderFunc(1, fn)
}

func (c *Canvas3d) AppendDefaultObj(obj gl.Canvas3D) {
	c.AppendObj(1, obj)
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
func (c *Canvas3d) Refresh() {
	c.Canvas3dObjs.Refresh()
}
func (c *Canvas3d) GetRenderObj() fyne.CanvasObject {
	return c.Canvas3dObjs
}
