package priorityqueue

type Orderable interface {
	OrderKey() int
}

type Tuple2 struct {
	Key   int
	First int
}

func (t Tuple2) OrderKey() int {
	return t.Key
}

type Tuple3 struct {
	Key    int
	First  int
	Second int
}

func (t Tuple3) OrderKey() int {
	return t.Key
}
