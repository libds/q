package q

import "errors"

var QEmpty = errors.New("Empty Queue")

type Q struct {
	qu []string
}

// Initiate new Q instance
func New() *Q {
	return &Q{qu: []string{}}
}

// Removes a job from queue and returns it.
// If there's no jobs then it will return QEmpty
func (q *Q) Deq() (value string, QEmpty error) {
	if len(q.qu) == 0 {
		return "", QEmpty
	}
	value = q.qu[0]
	q.qu = append(q.qu[:0], q.qu[1:]...)
	return value, nil
}

// Enqueue a job
func (q *Q) Enq(v string) {
	q.qu = append(q.qu, v)
}

// If queue is empty it'll return true
func (q *Q) Empty() bool {
	return len(q.qu) == 0
}

// Return the number of jobs in the queue
func (q *Q) Size() int {
	return len(q.qu)
}

// Returns the underlying array of the queue
func (q *Q) Arr() []string {
	return q.qu
}
