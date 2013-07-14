package planet

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})


func (s *S) TestRectArea(c *C) {
	r := Rect{Width: 10, Height: 5}
	c.Assert(r.Area(), Equals, 50)
}

func (s *S) TestRectPerim(c *C) {
	r := Rect{Width: 10, Height: 5}
	c.Assert(r.Perim(), Equals, 30)
}