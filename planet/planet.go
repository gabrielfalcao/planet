package planet

type Rect struct {
	Width int
	Height int
}

func (r *Rect) Area() int {
	return r.Width * r.Height
}

func (r *Rect) Perim() int {
	return 2*r.Width +  2*r.Height
}