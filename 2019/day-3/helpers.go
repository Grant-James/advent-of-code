package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getOptimizedIntersection(i string, i2 string) (int, error) {
	_, points, err := getDistance(i, i2)
	if err != nil {
		return 0, err
	}

	_, _, a, err := getLines(i)
	if err != nil {
	   return 0, err
	}

	_, _, a2, err := getLines(i2)
	if err != nil {
		return 0, err
	}

	var stepTo1 []int
	var stepTo2 []int
	var steps int
	for _, p := range points {
		steps := 0
		for _, line := range a {
			isInte, leftOver := isPIntersection(line, p, line.initial)
			if isInte {
				stepTo1 = append(stepTo1, steps + leftOver)
			}

			steps += line.val
		}

		steps = 0
		for _, line := range a2 {
			isInte, leftOver := isPIntersection(line, p, line.initial)
			if isInte {
				stepTo2 = append(stepTo2, steps + leftOver)
			}

			steps += line.val
		}
	}

	if len(stepTo1) != len(stepTo2) {
		return 0, fmt.Errorf("lens of intersections differ")

	}

	steps = 0
	for i, s := range stepTo1 {
		combined := s + stepTo2[i]
		if steps == 0 || combined < steps {
			steps = combined
		}
	}

	return steps, nil
}

func getDistance(i string, i2 string) (int, []*point, error) {
	h, v, _, err := getLines(i)
	if err != nil {
	    return 0, nil,err
	}
	
	h2, v2, _, err := getLines(i2)
	if err != nil {
	    return 0, nil, err
	}

	var points []*point
	for _, l := range h {
		for _, l2 := range v2 {
			point := getIntersection(l, l2)
			if point != nil {
				points = append(points, point)
			}
		}
	}
	for _, l := range h2 {
		for _, l2 := range v {
			point := getIntersection(l, l2)
			if point != nil {
				points = append(points, point)
			}
		}
	}

	dist := 0
	for _, p := range points {
		curDist := int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
		if dist == 0 || curDist < dist{
			dist = curDist
		}
	}

	return dist, points, nil
}

func getIntersection(a *line, b *line) *point {
	a.order()
	b.order()
	h := getH(a, b)
	v := getV(a, b)
	if h == nil || v == nil {
		return nil
	}

	if v.a.y <= h.a.y && v.b.y >= h.a.y && h.a.x <= v.a.x && h.b.x >= v.a.x {
		return &point{v.a.x, h.a.y}
	}
	return nil
}

func isPIntersection(l *line, p *point, initial string) (bool, int) {
	l.order()
	leftOver := 0
	if p.y == l.a.y {
		if p.x > l.a.x && p.x < l.b.x {
			if initial == "L" {
				leftOver = l.b.x - p.x
			} else if initial == "R" {
				leftOver = p.x - l.a.x
			}
			return true, int(math.Abs(float64(leftOver)))
		}
	} else if p.x == l.a.x {
		if p.y > l.a.y && p.y < l.b.y {
			if initial == "U" {
				leftOver = p.y - l.a.y
			} else if initial == "D" {
				leftOver = l.b.y - p.y
			}
			return true, int(math.Abs(float64(leftOver)))
		}
	}
	return false, 0
}

func getH(a *line, b *line) *line {
	if a.a.y == a.b.y {
		return a
	} else if b.a.y == b.b.y {
		return b
	}
	return nil
}

func getV(a *line, b *line) *line {
	if a.a.x == a.b.x {
		return a
	} else if b.a.x == b.b.x {
		return b
	}
	return nil
}

func getLines(data string) ([]*line, []*line, []*line, error) {
	var outA []*line
	var outH []*line
	var outV []*line
	ds := strings.Split(data, ",")
	newX := 0
	newY := 0
	curX := 0
	curY := 0
	for _, d := range ds {
		initial, val, err := getData(d)
		if err != nil {
		    return nil, nil, nil, err
		}

		switch initial {
			case "U":
				newY += val
			case "D":
				newY -= val
			case "L":
				newX -= val
			case "R":
				newX += val
		}

		var newPos *point
		curPos := &point{x: curX, y: curY}
		if initial == "U" || initial == "D" {
			newPos = &point{x: curX, y: newY}
			ol := &line{a: curPos, b: newPos, val: val, initial: initial}
			outV = append(outV, ol)
			outA = append(outA, ol)
		} else {
			newPos = &point{x: newX, y: curY}
			ol := &line{a: curPos, b: newPos, val: val, initial: initial}
			outA = append(outA, ol)
			outH = append(outH, ol)
		}
		curX = newX
		curY = newY
	}
	return outH, outV, outA, nil
}

func getData(d string) (string, int, error) {
	val, err := strconv.Atoi(d[1:])
	if err != nil {
	    return "", 0, err
	}

	return d[:1], val, nil
}