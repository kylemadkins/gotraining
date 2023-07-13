package oop

import (
	"fmt"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Invalid length of %v", length)
	}
	sq := Square{
		x,
		y,
		length,
	}
	return &sq, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}
