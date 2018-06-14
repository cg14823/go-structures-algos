package structures

import (
	"testing"
)

func Test_SingleLinkList_AddLast(t *testing.T) {
	l := CreateSingleLinkList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	secondNode := l.Head.Next
	if l.Head.Value != 1 {
		t.Errorf("the value of head should be equal to 1 instead %d\n", l.Head.Value)
	}

	if secondNode.Value != 2 {
		t.Errorf("the value should be equal to 2 instead %d\n", secondNode.Value)
	}

	if thirdNode := secondNode.Next; thirdNode.Value != 3 {
		t.Errorf("the value should be equal to 3 instead %d\n", thirdNode.Value)
	}
}

func Test_SingleLinkedList_AddFirst(t *testing.T) {
	l := CreateSingleLinkList()
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	secondNode := l.Head.Next
	if l.Head.Value != 3 {
		t.Errorf("the value of head should be equal to 3 instead %d\n", l.Head.Value)
	}

	if secondNode.Value != 2 {
		t.Errorf("the value should be equal to 2 instead %d\n", secondNode.Value)
	}

	if thirdNode := secondNode.Next; thirdNode.Value != 1 {
		t.Errorf("the value should be equal to 1 instead %d\n", thirdNode.Value)
	}
}

func Test_SingleLinkedList_AddAfter(t *testing.T) {
	l := CreateSingleLinkList()
	errOne := l.AddAfter(1, 2)
	l.AddFirst(1)
	errTwo := l.AddAfter(1, 2)
	errThree := l.AddAfter(1, 3)
	errFour := l.AddAfter(2, 4)

	secondNode := l.Head.Next
	thirdNode := secondNode.Next
	fourthNode := thirdNode.Next

	if errOne == nil {
		t.Errorf("using AddAfter() on an empty list should return an error instead it returned nil")
	}

	if secondNode.Value != 3 || errTwo != nil {
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", secondNode.Value, errTwo.Error())
	}

	if thirdNode.Value != 2 || errThree != nil {
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", thirdNode.Value, errThree.Error())
	}

	if fourthNode.Value != 4 || errFour != nil {
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", fourthNode.Value, errFour.Error())
	}
}

func Test_SingleLinkedList_AddBefore(t *testing.T) {
	l := CreateSingleLinkList()
	errOne := l.AddBefore(1, 2)
	l.AddFirst(1)
	errTwo := l.AddBefore(1, 2)
	errThree := l.AddBefore(1, 3)
	errFour := l.AddBefore(2, 4)
	secondNode := l.Head.Next
	thirdNode := secondNode.Next
	fourthNode := thirdNode.Next

	if errOne == nil {
		t.Errorf("using AddBefore() on an empty list should return an error instead it returned nil")
	}

	if l.Head.Value != 4 {
		t.Errorf("the value should be 4 instead %d\n", l.Head.Value)
	}
	if secondNode.Value != 2 || errTwo != nil {
		t.Errorf("the value should be equal to 2 instead %d\nError: %s", secondNode.Value, errTwo.Error())
	}

	if thirdNode.Value != 3 || errThree != nil {
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", thirdNode.Value, errThree.Error())
	}

	if fourthNode.Value != 1 || errFour != nil {
		t.Errorf("the value should be equal to 1 instead %d\nError: %s", fourthNode.Value, errFour.Error())
	}
}

func Test_SingleLinkedList_Delete(t *testing.T) {
	l := CreateSingleLinkList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	err0 := l.Delete(1)
	if l.Head.Value != 2 && err0 != nil {
		t.Errorf("the value should not be equal to 1\nError: %s", err0.Error())
	}

	err1 := l.Delete(23)
	if err1 == nil {
		t.Errorf("using Delete with a non-existing key should return an error")
	}
}

func Test_SingleLinkedList_Exists(t *testing.T) {
	l := CreateSingleLinkList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	if !l.Exists(3) {
		t.Errorf("should have return true")
	}

	if l.Exists(23) {
		t.Errorf("should have return false")
	}
}

func Test_SingleLinkedList_Iterator(t *testing.T) {
	l := CreateSingleLinkList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddFirst(5)
	l.AddLast(3)

	it := l.Iterator()
	numbers := []int{5, 1, 2, 3}
	ix := 0
	for it.HasNext() {
		val := it.Next()
		if val != numbers[ix] {
			t.Errorf("iterator does not work element should be %d but instead %d", numbers[ix], val)
		}
		ix++
	}

	if ix != 4 {
		t.Errorf("Should have iterate through 3 elements instead %d", ix)
	}
}
