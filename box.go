package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("the box goes out of the shapesCapacity range")
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New("index doesn't exist or index went out of the range")
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	result, err := b.GetByIndex(i)

	if err == nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	}

	return result, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	result, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}

	b.shapes[i] = shape

	return result, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0
	for _, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0
	for _, shape := range b.shapes {
		result += shape.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	oldLen := len(b.shapes)
	for i := oldLen - 1; i >= 0; i-- {
		if _, ok := b.shapes[i].(*Circle); !ok {
			continue
		}

		b.shapes[i] = b.shapes[0]
		b.shapes = b.shapes[1:]
	}

	if len(b.shapes) == oldLen {
		return errors.New("circles are not exist in the list")
	}

	return nil
}
