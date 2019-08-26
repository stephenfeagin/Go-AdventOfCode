package main

import (
	"errors"
	"sort"
)

// Point defines an x,y coordinate
type Point struct {
	x, y int
}

// GetDistance returns the Manhattan distance between points p and other
func (p Point) GetDistance(other Point) int {
	xDist := AbsInt(p.x - other.x)
	yDist := AbsInt(p.y - other.y)
	return xDist + yDist
}

// FindNearestPoint determines which point in a list of points is closest to point p. If there are
// more than one points tied for the nearest, return an error
func (p Point) FindNearestPoint(points []Point) (Point, error) {
	distanceMap := make(map[int][]Point)

	for _, pt := range points {
		dist := p.GetDistance(pt)
		distanceMap[dist] = append(distanceMap[dist], pt)
	}

	distances := make([]int, len(distanceMap))
	i := 0
	for d := range distanceMap {
		distances[i] = d
		i++
	}
	sort.Ints(distances)

	nearest := distanceMap[distances[0]]
	if len(nearest) > 1 {
		return Point{}, errors.New("More than one point tied for closest")
	}

	return nearest[0], nil
}

// AbsInt calculates the absolute value of an integer
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DefineCanvas returns all unique points in a plain bounded by the provided points
func DefineCanvas(points []Point) []Point {
	n := len(points)
	// Get all x and y coordinates and calculate the max and min for each
	x := make([]int, n)
	y := make([]int, n)
	for i, pt := range points {
		x[i] = pt.x
		y[i] = pt.y
	}
	sort.Ints(x)
	sort.Ints(y)
	xMin := x[0]
	xMax := x[n-1]
	yMin := y[0]
	yMax := y[n-1]

	// You can use a bool-valued map as a set by initializing an empty map then setting the
	// "present" values to true.
	canvas := make(map[Point]bool)
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			pt := Point{i, j}
			canvas[pt] = true
		}
	}

	results := make([]Point, len(canvas))
	i := 0
	for pt := range canvas {
		results[i] = pt
		i++
	}

	return results
}
