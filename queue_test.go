package goqueue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int]()
	for i := 0; i < 3; i++ {
		q.Add(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Get()
		if !ok {
			t.Fatalf("failed to retrieve %d", i)
		}
		if i != item {
			t.Fatalf("Got %d, expected %d", item, i)
		}
	}
	item, ok := q.Get()
	if ok {
		t.Fatalf("Got unexpected item %d from empty queue", item)
	}
}

func TestResize(t *testing.T) {
	q := New[int]()
	for i := 0; i < 16; i++ {
		q.Add(i)
	}

	if q.length != 16 {
		t.Fatalf("Bad length")
	}

	for i := 0; i < 16; i++ {
		item, ok := q.Get()
		if !ok {
			t.Fatalf("failed to retrieve %d", i)
		}
		if i != item {
			t.Fatalf("Got %d, expected %d", item, i)
		}
	}
	item, ok := q.Get()
	if ok {
		t.Fatalf("Got unexpected item %d from empty queue", item)
	}

}
