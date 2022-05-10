package mheap

type Orderable interface {
	OrderKey() int
}

type Tuple2 struct {
	First  int
	Second int
}

func (t Tuple2) OrderKey() int {
	return t.First
}

type Tuple3 struct {
	First  int
	Second int
	Third  int
}

func (t Tuple3) OrderKey() int {
	return t.First
}
