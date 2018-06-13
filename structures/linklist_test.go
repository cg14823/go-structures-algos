package structures

import "testing"

func TestSingleLinkListAddLast(t *testing.T){
	l := CreateSingleLinkList()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	secondNode := l.Head.Next
	if l.Head.Value != 1 {
		t.Errorf("the value of head should be equal to 1 instead %d\n", l.Head.Value)
	}

	if secondNode.Value!= 2 {
		t.Errorf("the value should be equal to 2 instead %d\n", secondNode.Value)
	}

	if thirdNode := secondNode.Next; thirdNode.Value != 3 {
		t.Errorf("the value should be equal to 3 instead %d\n", thirdNode.Value)
	}
}

func TestSingleLinkedListAddFirst(t *testing.T){
	l := CreateSingleLinkList()
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	secondNode := l.Head.Next
	if l.Head.Value != 3 {
		t.Errorf("the value of head should be equal to 3 instead %d\n", l.Head.Value)
	}

	if secondNode.Value!= 2 {
		t.Errorf("the value should be equal to 2 instead %d\n", secondNode.Value)
	}

	if thirdNode := secondNode.Next; thirdNode.Value != 1 {
		t.Errorf("the value should be equal to 1 instead %d\n", thirdNode.Value)
	}
}

func TestSingleLinkedListAddAfter(t *testing.T){
	l := CreateSingleLinkList()
	errOne := l.AddAfter(1, 2)
	l.AddFirst(1)
	errTwo := l.AddAfter(1,2)
	errThree := l.AddAfter(1,3)
	errFour := l.AddAfter(2,4)

	secondNode := l.Head.Next
	thirdNode := secondNode.Next
	fourthNode := thirdNode.Next

	if errOne == nil {
		t.Errorf("using AddAfter() on an empty list should return an error instead it returned nil")
	}

	if secondNode.Value != 3 || errTwo != nil{
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", secondNode.Value, errTwo.Error())
	}

	if thirdNode.Value != 2  || errThree != nil{
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", thirdNode.Value, errThree.Error())
	}

	if fourthNode.Value != 4  || errFour != nil{
		t.Errorf("the value should be equal to 3 instead %d\nError: %s", fourthNode.Value, errFour.Error())
	}
}