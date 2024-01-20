package camera

import (
	"github.com/qiwikgames/babos2d/pkg/engine/gameapp/position"
)

type Camera struct {
	Pos    position.Position[float64]
	Width  float64
	Height float64
}
