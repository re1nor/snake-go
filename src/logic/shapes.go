package snake

type Point struct {
	X, Y float64
}

const (
	Top    = iota
	Right  = iota
	Bottom = iota
	Left   = iota
)

type Dir int

func (d Dir) Exec(p Point) Point {
	switch d {
	case Top:
		return Point{X: p.X, Y: p.Y + 1}
	case Right:
		return Point{X: p.X + 1, Y: p.Y}
	case Bottom:
		return Point{X: p.X, Y: p.Y - 1}
	case Left:
		return Point{X: p.X - 1, Y: p.Y}
	}
}
