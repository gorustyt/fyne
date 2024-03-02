package canvas3d

import (
	"fmt"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
	"image"
	"os"
)

var _ gl.Canvas3D = (*Texture)(nil)

type pathTexture string

func (p pathTexture) Image() image.Image {
	f, err := os.Open(string(p))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}

type imgTexture struct {
	img image.Image
}

func (p imgTexture) Image() image.Image {
	return p.img
}

type itexture interface {
	Image() image.Image
}
type Texture struct {
	paths       []itexture
	tex         []context.Texture
	customAttrs []string
	MixParams   float32
	imgs        map[int]image.Image
	levelData   map[int]map[int32][]uint8
}

func (tex *Texture) NeedShader() bool {
	return true
}

func (tex *Texture) InitOnce(p *gl.Painter3D) {
	tex.createTexture(p)
}

func NewTexture() *Texture {
	return &Texture{
		levelData: map[int]map[int32][]uint8{},
		MixParams: 0.2,
		imgs:      make(map[int]image.Image),
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

func (tex *Texture) AppendPathWithCustomAttr(attr string, p string) {
	tex.AppendPath(p)
	tex.customAttrs = append(tex.customAttrs, attr)
}

func (tex *Texture) AppendImageWithCustomAttr(attr string, p image.Image) {
	tex.AppendImage(p)
	tex.customAttrs = append(tex.customAttrs, attr)
}

func (tex *Texture) AppendImage(p image.Image) {
	tex.paths = append(tex.paths, &imgTexture{
		img: p,
	})
}

func (tex *Texture) AppendPath(p string) {
	tex.paths = append(tex.paths, pathTexture(p))
}
func (tex *Texture) SetDefaultLevelData(level int32, data []uint8) {
	tex.SetLevelData(0, level, data)
}

func (tex *Texture) SetLevelData(index int, level int32, data []uint8) {
	v, ok := tex.levelData[index]
	if ok {
		v = map[int32][]uint8{}
		tex.levelData[index] = v
	}
	v[level] = data
}

func (tex *Texture) createTexture(painter *gl.Painter3D) {
	for i, v := range tex.paths {
		if i < len(tex.tex) {
			continue
		}
		tex.tex = append(tex.tex, painter.MakeTexture(v.Image(), gl.GetTextureByIndex(i), tex.levelData[i]))
	}
}
