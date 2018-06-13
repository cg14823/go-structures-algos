package structures

import "testing"

// STACK  TESTS -------------------------------------------------------------------
func TestStackPopAndPush1(t *testing.T){
	s := InitStack()
	s.Push(1)
	s.Push(2)

	if thisIs2, _  := s.Pop(); thisIs2 != 2{
		t.Errorf("Stack Pop did not work expected 2 got %d\n", thisIs2)
	}
	if thisIs1, _  := s.Pop(); thisIs1 != 1{
		t.Errorf("Stack Pop did not work expected 1 got %d\n", thisIs1)
	}
}

func TestStackPopAndPush2(t *testing.T){
	s := InitStack()
	_ , err := s.Pop()
	if err == nil {
		t.Errorf("Pop empty stack should return error")
	}
}

func TestStackSize(t *testing.T) {
	s := InitStack()
	if size := s.Size(); size != 0 {
		t.Errorf("Size should be 0 instead %d\n", size)
	}

	s.Push(1)
	s.Push(2)
	if size := s.Size(); size != 2 {
		t.Errorf("Size should be 2 instead %d\n", size)
	}
	s.Pop()
	if size := s.Size(); size != 1 {
		t.Errorf("Size should be 1 instead %d\n", size)
	}
}

// QUEUE TEST -------------------------------------------------------------------------
func TestQueuePopAndPush1(t *testing.T){
	s := InitQueue()
	s.Push(1)
	s.Push(2)

	if thisIs1, _  := s.Pop(); thisIs1 != 1{
		t.Errorf("Queue Pop did not work expected 1 got %d\n", thisIs1)
	}
	if thisIs2, _  := s.Pop(); thisIs2 != 2{
		t.Errorf("Queue Pop did not work expected 2 got %d\n", thisIs2)
	}
}

func TestQueuePopAndPush2(t *testing.T){
	s := InitQueue()
	_ , err := s.Pop()
	if err == nil {
		t.Errorf("Pop empty queue should return error")
	}
}

func TestQueueSize(t *testing.T) {
	s := InitQueue()
	if size := s.Size(); size != 0 {
		t.Errorf("Size should be 0 instead %d\n", size)
	}

	s.Push(1)
	s.Push(2)
	if size := s.Size(); size != 2 {
		t.Errorf("Size should be 2 instead %d\n", size)
	}
	s.Pop()
	if size := s.Size(); size != 1 {
		t.Errorf("Size should be 1 instead %d\n", size)
	}
}