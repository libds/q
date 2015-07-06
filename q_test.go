package q

import "testing"

const job1 string = "Job1"
const job2 string = "Job2"

func TestQ(t *testing.T) {
	qu := New()
	qu.Enq(job1)
	qu.Enq(job2)

	if v := qu.Size(); v != 2 {
		t.Fatal("Q size must be 2")
	}

	if v := qu.Empty(); v {
		t.Fatal("Q must not be empty")
	}

	job, err := qu.Deq()

	if err != nil {
		t.Fatal("Q must not return error")
	}

	if job != job1 {
		t.Fatal("job must equal job1")
	}

	job, err = qu.Deq()

	if err != nil {
		t.Fatal("Q must not return error")
	}

	if job != job2 {
		t.Fatal("job must equal job2")
	}

	_, err = qu.Deq()
	if err == QEmpty {
		t.Fatal("Q must return QEmpty")
	}

	if qu.Size() != 0 {
		t.Fatal("Q size must be 0")
	}

	if !qu.Empty() {
		t.Fatal("Q size must be empty")
	}

	if len(qu.Arr()) != 0 {
		t.Fatal("Q array length must equal 0")
	}

}
