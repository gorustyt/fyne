package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/app"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/canvas3d_render"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	points := canvas3d.NewPoints().SetPointWidth(10)
	point := canvas3d.NewPoint()
	point.Colorubs = []uint8{255, 0, 0, 255}
	point.Pos = &mgl32.Vec3{-0.2, 0, -1.5}
	points.AppendPoint(point)

	point = canvas3d.NewPoint()
	point.Colorubs = []uint8{0, 255, 0, 255}
	point.Pos = &mgl32.Vec3{-0.1, 0, -1.5}
	points.AppendPoint(point)

	point.Colorubs = []uint8{0, 0, 255, 255}
	point.Pos = &mgl32.Vec3{-0.0, 0, -1.5}
	points.AppendPoint(point)
	c := canvas3d_render.NewCanvas3d()
	c.AppendDefaultObj(points)
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		c,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome 😀")
		}),
	))

	w.ShowAndRun()
}
