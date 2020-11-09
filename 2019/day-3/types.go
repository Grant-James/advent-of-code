package main

type line struct {
	a *point
	b *point
	val int
	initial string
}

type point struct {
	x int
	y int
}

func (l *line) order() {
	if l.a.x > l.b.x {
		point := &line{
			a: &point{l.b.x, l.a.y},
			b: &point{l.a.x, l.b.y},
		}
		l.a = point.a
		l.b = point.b
	}

	if l.a.y > l.b.y {
		point := &line{
			a: &point{l.a.x, l.b.y},
			b: &point{l.b.x, l.a.y},
		}
		l.a = point.a
		l.b = point.b
	}
}