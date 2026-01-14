package queue

type Queue[T any] []T

func (self *Queue[T]) Push(x T) {
	*self = append(*self, x)
}

func (self *Queue[T]) Pop() (T, bool) {
	h := *self
	var out T
	if len(*self) == 0 {
		return out, false
	}
	l := len(h)
	out, *self = h[0], h[1:l]

	return out, true
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}