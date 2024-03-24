package shake

import (
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

const (
	GameW = 720.0
	GameH = 720.0
)

type Game struct {
	cv  *canvas.Canvas
	wnd *sdlcanvas.Window
}

func NewGame() *Game {
	wnd, cv, err := sdlcanvas.CreateWindow(1080, 750, "hello, Snake!")
	if err != nil {
		panic(err)
	}
	g := &Game{
		cv:  cv,
		wnd: wnd,
	}
	return g
}

func (g *Game) Run() {
	g.renderLoop()
}

func (g *Game) renderLoop() {
	gameAreaSP := Point{15, 15}
	gameAreaEP := Point{15 + GameW, 15 + GameH}

	//cellW := GameW / 20
	//cellH := GameH / 20

	g.wnd.MainLoop(func() {

		//clear
		g.cv.ClearRect(0, 0, 1080, 750)

		//отрисовывать  мир
		g.cv.BeginPath()
		g.cv.SetFillStyle("#333")                                                   //цвет чего-то там
		g.cv.FillRect(gameAreaSP.x, gameAreaSP.y, gameAreaEP.x-15, gameAreaEP.y-15) //цвет фона
		g.cv.Stroke()

		//render snake
		//render score
	})
}
