package canvas3d

import (
	"fmt"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
	"image"
	"os"
)

var _ gl.Canvas3D = (*Texture)(nil)

type Texture struct {
	paths       []string
	tex         []context.Texture
	customAttrs []string
	MixParams   float32
}

func (tex *Texture) InitOnce(p *gl.Painter3D) {
	tex.createTexture(p)
}

func NewTexture() *Texture {
	return &Texture{
		MixParams: 0.2,
	}
}
func (tex *Texture) Init(p *gl.Painter3D) {
	if len(tex.customAttrs) > 0 {
		for i, v := range tex.customAttrs {
			p.Uniform1i(v, int32(i))
			p.ActiveTexture(gl.GetTextureByIndex(i))
			p.BindTexture(tex.tex[i])
		}
	} else {
		for i, v := range tex.tex {
			p.Uniform1i(fmt.Sprintf("texture%v", i+1), int32(i))
			p.ActiveTexture(gl.GetTextureByIndex(i))
			p.BindTexture(v)
		}
	}
	p.Uniform1f("mixParams", tex.MixParams)
}

func (tex *Texture) After(p *gl.Painter3D) {

}
func (tex *Texture) AppendCustomAttr(attr string) {
	tex.customAttrs = append(tex.customAttrs, attr)
}
func (tex *Texture) AppendPath(p string) {
	tex.paths = append(tex.paths, p)
}

func (tex *Texture) createTexture(painter *gl.Painter3D) {
	openFile := func(p string, index int) {
		f, err := os.Open(p)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		tex.tex = append(tex.tex, painter.MakeTexture(img, gl.GetTextureByIndex(index)))
	}
	for i, v := range tex.paths {
		if i < len(tex.tex) {
			continue
		}
		openFile(v, i)
	}
}
