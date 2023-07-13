package oop

import (
	"fmt"
	"io"
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

type Capper struct {
	Wtr io.Writer
}

func (c *Capper) Write(p []byte) (int, error) {
	diff := byte('a' - 'A')
	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.Wtr.Write(out)
}
