package shake

type Point struct {
	x, y float64
}

type Dir int

const (
	Top    = iota
	Right  = iota
	Left   = iota
	Bottom = iota
)

func (d Dir) Exec(p Point) Point {
	//принимает точку - двигает на новую позицию - возвр новую точку
	switch d {
	case Top:
		return Point{p.x, p.y + 1}
	case Right:
		return Point{p.x + 1, p.y}
	case Left:
		return Point{p.x - 1, p.y}
	case Bottom:
		return Point{p.x, p.y - 1}
	}
	return Point{-1, -1}
}
