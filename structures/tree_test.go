package structures

import (
	"testing"
)

func Test_BST_Insert(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(2)
	tr.Insert(3)
	tr.Insert(1)

	leftNode := tr.Root.Left
	rightNode := tr.Root.Right

	if tr.Root.Value != 2 {
		t.Errorf("Root should be 2 instrad %d\n", tr.Root.Value)
	}
	if leftNode.Value != 1 {
		t.Errorf("Root should be 1 instrad %d\n", leftNode.Value)

	}
	if rightNode.Value != 3 {
		t.Errorf("Root should be 3 instrad %d\n", rightNode.Value)

	}
}

func Test_BST_Contains(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(2)
	tr.Insert(3)
	tr.Insert(1)

	if !tr.Contains(1) {
		t.Errorf("Should have return true\n")
	}
	if !tr.Contains(2) {
		t.Errorf("Should have return true\n")
	}
	if !tr.Contains(3) {
		t.Errorf("Should have return true\n")
	}
	if tr.Contains(11) {
		t.Errorf("Should have return false\n")
	}
}

func Test_BST_DeleteRoot(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)

	tr.Delete(20)
	leftNode := tr.Root.Left
	rightNode := tr.Root.Right
	leftRightNode := leftNode.Right
	leftLeftNode := leftNode.Left

	if tr.Root.Value != 16 {
		t.Errorf("Head should be 16 but instead %d\n", tr.Root.Value)
	}
	if leftNode.Value != 15 {
		t.Errorf("left node should be 15 but instead %d\n", leftNode.Value)

	}
	if rightNode.Value != 25 {
		t.Errorf("left node should be 25 but instead %d\n", rightNode.Value)
	}
	if leftRightNode != nil {
		t.Errorf("left right node should be nil but instead %v\n", leftRightNode)
	}
	if leftLeftNode.Value != 11 {
		t.Errorf("left left node should be 11 but instead %d\n", leftLeftNode.Value)
	}
}

func Test_BST_DeleteLeaf(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)

	tr.Delete(16)
	tr.Delete(22)

	leftNode := tr.Root.Left
	rightNode := tr.Root.Right

	if tr.Root.Value != 20 {
		t.Errorf("Head should be 16 but instead %d\n", tr.Root.Value)
	}
	if leftNode.Value != 15 {
		t.Errorf("left node should be 15 but instead %d\n", leftNode.Value)

	}
	if rightNode.Value != 25 {
		t.Errorf("left node should be 25 but instead %d\n", rightNode.Value)
	}
	if leftNode.Right != nil {
		t.Errorf("left right node should be nil but instead %v\n", leftNode.Right)
	}
	if rightNode.Left != nil {
		t.Errorf("right left node should be nil but instead %v\n", rightNode.Left)
	}
}

func Test_BST_DeleteWithTwoChildren(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)

	tr.Delete(15)
	tr.Delete(25)

	leftNode := tr.Root.Left
	rightNode := tr.Root.Right

	if tr.Root.Value != 20 {
		t.Errorf("Head should be 20 but instead %d\n", tr.Root.Value)
	}
	if leftNode.Value != 11 {
		t.Errorf("left node should be 11 but instead %d\n", leftNode.Value)
	}
	if rightNode.Value != 22 {
		t.Errorf("left node should be 22 but instead %d\n", rightNode.Value)
	}
}

func Test_BST_Delete_NotExist(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)

	deleted := tr.Delete(27)

	if deleted {
		t.Errorf("Should have returned false")
	}
}

func Test_BST_IT_InOrder(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)
	tr.Insert(30)

	it := tr.Iterator("in")
	inOrder := []int{11, 15, 16, 20, 22, 25, 28, 30}
	ix := 0
	for it.HasNext() {
		v := it.Next()
		if v != inOrder[ix] {
			t.Errorf("Value should have been %d instead %d\n", inOrder[ix], v)
		}
		ix++
	}

	if ix != len(inOrder) {
		t.Errorf("Should have iterated over %d values instead %d", len(inOrder), ix)
	}

}

func Test_BST_IT_PreOrder(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)
	tr.Insert(30)

	it := tr.Iterator("pre")
	inOrder := []int{20, 15, 11, 16, 25, 22, 28, 30}
	ix := 0
	for it.HasNext() {
		v := it.Next()
		if v != inOrder[ix] {
			t.Errorf("Value should have been %d instead %d\n", inOrder[ix], v)
		}
		ix++
	}

	if ix != len(inOrder) {
		t.Errorf("Should have iterated over %d values instead %d", len(inOrder), ix)
	}

}

func Test_BST_IT_PostOrder(t *testing.T) {
	tr := CreateBinarySearchTree()
	tr.Insert(20)
	tr.Insert(15)
	tr.Insert(25)
	tr.Insert(11)
	tr.Insert(16)
	tr.Insert(22)
	tr.Insert(28)
	tr.Insert(30)

	it := tr.Iterator("post")
	inOrder := []int{11, 16, 15, 22, 30, 28, 25, 20}
	ix := 0
	for it.HasNext() {
		v := it.Next()
		if v != inOrder[ix] {
			t.Errorf("Value should have been %d instead %d\n", inOrder[ix], v)
		}
		ix++
	}

	if ix != len(inOrder) {
		t.Errorf("Should have iterated over %d values instead %d", len(inOrder), ix)
	}

}
