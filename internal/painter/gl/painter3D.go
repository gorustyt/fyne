package gl

import "fyne.io/fyne/v2"

type Painter3D struct {
	prog Program //sharder
	context
}

func NewPainter3D(ctx context) *Painter3D {
	return &Painter3D{context: ctx}
}

func (p *Painter3D) Program() Program {
	return p.Program()
}

func (p *Painter3D) HasShader() bool {
	return p.prog == 0
}

func (p *Painter3D) DefineVertexArray(name string, size, stride, offset int) {
	vertAttrib := p.GetAttribLocation(p.prog, name)
	p.context.EnableVertexAttribArray(vertAttrib)
	p.VertexAttribPointerWithOffset(vertAttrib, size, float, false, stride*floatSize, offset*floatSize)
}

func (p *Painter3D) BindTexture(texture Texture) {
	p.context.BindTexture(texture2D, texture)
}

type Canvas3D interface {
	Init(p *Painter3D)
	Draw(p *Painter3D, pos fyne.Position, Frame fyne.Size)
	After(p *Painter3D)
}

type Canvas3dObj struct {
	Painter *Painter3D
	objs    []Canvas3D

	vertStr, fragStr string
}

func (c *Canvas3dObj) SetShaderConfig(vertStr, fragStr string) {
	c.vertStr, c.fragStr = vertStr, fragStr
}

func (c *Canvas3dObj) Move(position fyne.Position) {

}

func (c *Canvas3dObj) Position() fyne.Position {
	return fyne.Position{}
}

func (c *Canvas3dObj) Hide() {

}

func (c *Canvas3dObj) Visible() bool {
	return true
}

func (c *Canvas3dObj) Show() {

}

func (c *Canvas3dObj) MinSize() fyne.Size {
	return fyne.Size{Width: 600, Height: 600}
}

func (c *Canvas3dObj) Resize(size fyne.Size) {

}

func (c *Canvas3dObj) Size() fyne.Size {
	return fyne.Size{Width: 600, Height: 600}
}

func (c *Canvas3dObj) Refresh() {

}

func NewCustomObj() fyne.CanvasObject {
	return &Canvas3dObj{}
}
