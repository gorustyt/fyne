package canvas3d

import "github.com/go-gl/mathgl/mgl32"

type Coordinate struct {
	Project mgl32.Mat4
	View    mgl32.Mat4
	Model   mgl32.Mat4
}
