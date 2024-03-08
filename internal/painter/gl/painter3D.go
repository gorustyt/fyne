package gl

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"strings"
)

type Painter3D struct {
	prog context.Program //sharder
	context.Context
	mode int
}

func NewPainter3D(ctx context.Context) *Painter3D {

	return &Painter3D{Context: ctx}
}
func (p *Painter3D) GetContext() context.Context {
	return p.Context
}
func (p *Painter3D) HasInit() bool {
	return p.prog != 0
}

func (p *Painter3D) DrawTrianglesByElement(index []uint32) {
	p.Context.DrawElementsArrays(Triangles, index)
}

func (p *Painter3D) DrawTriangles(count int) {
	p.Context.DrawArrays(Triangles, 0, count)
}

func (p *Painter3D) DefineVertexArray(name string, size, stride, offset int) {
	vertAttrib := p.GetAttribLocation(p.prog, name)
	p.Context.EnableVertexAttribArray(vertAttrib)
	p.VertexAttribPointerWithOffset(vertAttrib, size, float, false, stride*floatSize, offset*floatSize)
}

func (p *Painter3D) BindTexture(texture context.Texture) {
	p.Context.BindTexture(Texture2D, texture)
}

func (p *Painter3D) UniformMatrix4fv(name string, mat4 mgl32.Mat4) {
	p.Context.UniformMatrix4fv(p.prog, name, mat4)
}

func (p *Painter3D) Uniform1i(name string, v0 int32) {
	p.Context.Uniform1i(p.prog, name, v0)
}

func (p *Painter3D) Uniform1f(name string, v float32) {
	p.Context.Uniform1f(p.GetUniformLocation(p.prog, name), v)
}

func (p *Painter3D) Uniform2f(name string, v0, v1 float32) {
	p.Context.Uniform2f(p.GetUniformLocation(p.prog, name), v0, v1)
}

func (p *Painter3D) UniformVec3(name string, vec3 mgl32.Vec3) {
	p.Context.Uniform3f(p.GetUniformLocation(p.prog, name), vec3)
}

func (p *Painter3D) Uniform3f(name string, v0, v1, v2 float32) {
	p.Context.Uniform3f(p.GetUniformLocation(p.prog, name), mgl32.Vec3{v0, v1, v2})
}

func (p *Painter3D) Uniform4f(name string, v0, v1, v2, v3 float32) {
	p.Context.Uniform4f(p.GetUniformLocation(p.prog, name), v0, v1, v2, v3)
}

type Canvas3D interface {
	InitOnce(p *Painter3D)
	Init(p *Painter3D)
	After(p *Painter3D)
	NeedShader() bool
}

type Canvas3DBeforePainter interface {
	BeforeDraw(p context.Painter, pos fyne.Position, Frame fyne.Size)
}

type Canvas3DPainter interface {
	Draw(p *Painter3D, pos fyne.Position, Frame fyne.Size)
}

type Canvas3dObj struct {
	Painter          *Painter3D
	Objs             []Canvas3D
	RenderFuncs      []func(ctx context.Painter)
	vertStr, fragStr string
	proCache         map[string]context.Program
}

func (c *Canvas3dObj) SplitByNeedShader() (needs []Canvas3D, notNeeds []Canvas3D) {
	for _, v := range c.Objs {
		if v.NeedShader() {
			needs = append(needs, v)
		} else {
			notNeeds = append(notNeeds, v)
		}
	}
	return
}

func (c *Canvas3dObj) ChangeShader(vertStr, fragStr string) {
	if c.vertStr == vertStr || c.fragStr == fragStr {
		return
	}
	oldId := strings.Join([]string{c.vertStr, c.fragStr}, ",")
	var oldPro context.Program
	if c.Painter != nil {
		oldPro = c.Painter.prog
	}
	c.vertStr = vertStr
	c.fragStr = fragStr
	if oldId != "," {
		id := strings.Join([]string{c.vertStr, c.fragStr}, ",") //TODO 压缩字符串
		if pro, ok := c.proCache[id]; ok {
			c.Painter.prog = pro
		} else {
			c.Painter.prog = 0
		}
		c.proCache[oldId] = oldPro
	}
}

func (c *Canvas3dObj) GetShader() (vertStr, fragStr string) {
	return c.vertStr, c.fragStr
}

func (c *Canvas3dObj) InitOnce() {
	for _, v := range c.Objs {
		v.InitOnce(c.Painter)
	}
}

func (c *Canvas3dObj) Init() {
	c.Painter.EnableDepthTest()
	for _, v := range c.Objs {
		v.Init(c.Painter)
	}
	for _, v := range c.RenderFuncs {
		v(c.Painter)
	}
}

func (c *Canvas3dObj) BeforeDraw(pos fyne.Position, frame fyne.Size) {
	for _, v := range c.Objs {
		if cc, ok := v.(Canvas3DBeforePainter); ok {
			cc.BeforeDraw(c.Painter, pos, frame)
		}
	}
}

func (c *Canvas3dObj) Draw(pos fyne.Position, frame fyne.Size) {
	for _, v := range c.Objs {
		if cc, ok := v.(Canvas3DPainter); ok {
			cc.Draw(c.Painter, pos, frame)
		}
	}
}

func (c *Canvas3dObj) After() {
	for i := len(c.Objs) - 1; i >= 0; i-- {
		c.Objs[i].After(c.Painter)
	}
	c.Painter.DisableDepthTest()
}

type Canvas3dObjs struct {
	objs []*Canvas3dObj
}

func (c *Canvas3dObjs) GetCanvas3dObj(index int) *Canvas3dObj {
	return c.objs[index]
}

func (c *Canvas3dObjs) RangeCanvas3dObj(fn func(obj *Canvas3dObj) (stop bool)) {
	for _, v := range c.objs {
		if fn(v) {
			return
		}
	}
}

func (c *Canvas3dObjs) Dragged(ev *fyne.DragEvent) {
	c.RangeCanvas3dObj(func(obj *Canvas3dObj) (stop bool) {
		for _, v := range obj.Objs {
			if p, ok := v.(fyne.Draggable); ok {
				p.Dragged(ev)
			}
		}
		return false
	})

}
func (c *Canvas3dObjs) DragEnd() {
	c.RangeCanvas3dObj(func(obj *Canvas3dObj) (stop bool) {
		for _, v := range obj.Objs {
			if p, ok := v.(fyne.Draggable); ok {
				p.DragEnd()
			}
		}
		c.Refresh()
		return false
	})
}
func (c *Canvas3dObjs) Scrolled(ev *fyne.ScrollEvent) {
	c.RangeCanvas3dObj(func(obj *Canvas3dObj) (stop bool) {

		for _, v := range obj.Objs {
			if p, ok := v.(fyne.Scrollable); ok {
				p.Scrolled(ev)
			}
		}
		c.Refresh()
		return false
	})
}
func (c *Canvas3dObjs) Move(position fyne.Position) {

}

func (c *Canvas3dObjs) Position() fyne.Position {
	return fyne.Position{}
}

func (c *Canvas3dObjs) Hide() {

}

func (c *Canvas3dObjs) Visible() bool {
	return true
}

func (c *Canvas3dObjs) Show() {

}

func (c *Canvas3dObjs) MinSize() fyne.Size {
	return fyne.Size{Width: 600, Height: 1080}
}

func (c *Canvas3dObjs) Resize(size fyne.Size) {

}

func (c *Canvas3dObjs) Size() fyne.Size {
	return fyne.Size{Width: 600, Height: 1080}
}

func (c *Canvas3dObjs) Refresh() {
	canvas.Refresh(c)
}

func NewCustomObj() *Canvas3dObj {
	return &Canvas3dObj{proCache: map[string]context.Program{}}
}

func NewCustomObjs(n int) *Canvas3dObjs {
	if n <= 0 {
		n = 1
	}
	obj := &Canvas3dObjs{objs: make([]*Canvas3dObj, n)}
	for i := range obj.objs {
		obj.objs[i] = NewCustomObj()
	}
	return obj
}
