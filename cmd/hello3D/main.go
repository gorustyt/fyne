package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/app"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/canvas3d_render"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	points := canvas3d.NewPoints().SetPointWidth(30)
	point := canvas3d.NewPoint()
	point.Colorubs = []uint8{255, 0, 0, 255}
	point.Pos = &mgl32.Vec3{0.5, 0, 0}
	points.AppendPoint(point)

	point = canvas3d.NewPoint()
	point.Colorubs = []uint8{0, 255, 0, 255}
	point.Pos = &mgl32.Vec3{-0.5, 0, 0}
	points.AppendPoint(point)

	point = canvas3d.NewPoint()
	point.Colorubs = []uint8{0, 0, 255, 255}
	point.Pos = &mgl32.Vec3{-0.0, 0, 0.5}
	points.AppendPoint(point)

	lines := canvas3d.NewLines().SetLineWidth(30)
	lines.SetDrawLoop()
	lines.AppendLine(&canvas3d.Line{
		P1: canvas3d.Point{
			Pos:      &mgl32.Vec3{0, 0.5, 0.5},
			Colorubs: []uint8{255, 0, 0, 255},
		},
		P2: canvas3d.Point{
			Pos:      &mgl32.Vec3{0.0, 0.0, 0.5},
			Colorubs: []uint8{0, 255, 0, 255},
		},
	})

	lines.AppendLine(&canvas3d.Line{
		P2: canvas3d.Point{
			Pos:      &mgl32.Vec3{0.5, 0.0, -0.5},
			Colorubs: []uint8{0, 0, 255, 255},
		},
		P1: canvas3d.Point{
			//Pos:      &mgl32.Vec3{0, 0.0, -0.5},
			Colorubs: []uint8{0, 255, 255, 255},
		},
	})
	c := canvas3d_render.NewCanvas3d()
	c.AppendDefaultObj(points)
	c.AppendDefaultObj(lines)
	w.SetContent(c.GetRenderObj())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
