package walker

import (
	"math/rand"

	"github.com/fzipp/canvas"
)

type Walker struct {
	x, y   float64
	px, py float64
}

func NewWalker(canvasWidth, canvasHeight int) *Walker {
	pointX, pointY := float64(canvasWidth/2), float64(canvasHeight/2)
	return &Walker{pointX, pointY, pointX, pointY}
}

func (w *Walker) Draw(ctx *canvas.Context) {
	w.attractionMove()
	ctx.Rect(w.x, w.y, 1, 1)
	ctx.Fill()
}
func (w *Walker) Handle(evt canvas.Event) {
	e, ok := evt.(canvas.MouseMoveEvent)
	if !ok {
		return
	}
	w.px = float64(e.X)
	w.py = float64(e.Y)
}

func (w *Walker) attractionMove() {
	if w.px == w.x && w.py == w.y {
		newX := rand.Intn(3) - 1
		newY := rand.Intn(3) - 1
		w.x += float64(newX)
		w.y += float64(newY)
	} else {
		randX, randY := rand.Float32(), rand.Float32()
		newX, newY := -1+randX*2, -1+randY*2
		distX, distY := w.stdDevDistance(0.5, 1, newX), w.stdDevDistance(0.5, 1, newY)
		if w.x < w.px {
			if randX < 0.2 {
				w.x -= distX
			} else if randX > 0.5 {
				w.x += distX
			}
		} else {
			if randX < 0.2 {
				w.x += distX
			} else if randX > 0.5 {
				w.x -= distX
			}
		}

		if w.y < w.py {
			if randY < 0.2 {
				w.y -= distY
			} else if randY > 0.5 {
				w.y += distY
			}
		} else {
			if randY < 0.2 {
				w.y += distY
			} else if randY > 0.5 {
				w.y -= distY
			}
		}
	}
}

func (w *Walker) move() {
	newX := rand.Intn(3) - 1
	newY := rand.Intn(3) - 1
	w.x += float64(newX)
	w.y += float64(newY)
}

func (w *Walker) stdDevDistance(sd, mean, num float32) float64 {
	return float64(sd*num + mean)
}

func (w *Walker) monteCarloRand(min, max float64) float64 {
	for {
		probability, r2 := floatInBetween(min, max), floatInBetween(min, max)
		if r2 < probability {
			return r2
		}
	}
}

func floatInBetween(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
