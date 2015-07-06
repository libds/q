package q

import "errors"

var QEmpty = errors.New("Empty Queue")

type Q struct {
	qu []string
}

func New() *Q {
	return &Q{qu: []string{}}
}

func (q *Q) Deq() (value string, QEmpty error) {
	if len(q.qu) == 0 {
		return "", QEmpty
	}
	value = q.qu[0]
	q.qu = append(q.qu[:0], q.qu[1:]...)
	return value, nil
}

func (q *Q) Enq(v string) {
	q.qu = append(q.qu, v)
}

func (q *Q) Empty() bool {
	return len(q.qu) == 0
}

func (q *Q) Size() int {
	return len(q.qu)
}

func (q *Q) Arr() []string {
	return q.qu
}
