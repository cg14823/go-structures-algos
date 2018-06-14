package structures

type Iterator interface {
	HasNext() bool
	Next() int
}
type node struct {
	Value int
	Left  *node
	Right *node
}

type InOrderIterartor struct {
	s       []*node
	visited map[int]*node
}

func (bi *InOrderIterartor) HasNext() bool {
	return len(bi.s) != 0
}

func (bi *InOrderIterartor) Next() int {
	curr := bi.s[len(bi.s)-1]
	// Go down left tree first
	if l := curr.Left; l != nil {
		if _, ok := bi.visited[l.Value]; !ok {
			bi.s = append(bi.s, l)
			return bi.Next()
		}
	}
	// else pop root and move right tree
	el := curr.Value
	//add the current node to the visited ones
	bi.visited[el] = curr
	// remove the element from the stack
	bi.s = bi.s[:len(bi.s)-1]
	//if it has a right root add it to stack
	if r := curr.Right; r != nil {
		if _, ok := bi.visited[r.Value]; ok {
			bi.s = bi.s[:len(bi.s)-1]
		} else {
			bi.s = append(bi.s, r)
		}
	}
	return el
}

type PreOrderIterartor struct {
	s []*node
}

func (bi *PreOrderIterartor) HasNext() bool {
	return len(bi.s) != 0
}

func (bi *PreOrderIterartor) Next() int {
	curr := bi.s[len(bi.s)-1]
	bi.s = bi.s[:len(bi.s)-1]
	if curr.Right != nil {
		bi.s = append(bi.s, curr.Right)
	}
	if curr.Left != nil {
		bi.s = append(bi.s, curr.Left)
	}
	return curr.Value
}

type PostOrderIterartor struct {
	s       []*node
	visited map[int]*node
}

func (bi *PostOrderIterartor) HasNext() bool {
	return len(bi.s) != 0
}

func (bi *PostOrderIterartor) Next() int {
	curr := bi.s[len(bi.s)-1]
	if l := curr.Left; l != nil {
		if _, ok := bi.visited[l.Value]; !ok {
			bi.s = append(bi.s, l)
			return bi.Next()
		}
	}
	if r := curr.Right; r != nil {
		if _, ok := bi.visited[r.Value]; !ok {
			bi.s = append(bi.s, r)
			return bi.Next()
		}
	}
	el := curr.Value
	bi.visited[el] = curr
	bi.s = bi.s[:len(bi.s)-1]
	return el
}

// BinarySearchTree is a data structure used to store integers in a sorted manner.
// Nodes to the left are smaller than the current node and node to the right are bigger.
// No duplicates accepted
type BinarySearchTree struct {
	Root *node
	size int
}

func CreateBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
		size: 0,
	}
}

// Insert elements into tree. If returns true it inserted it else not.
func (t *BinarySearchTree) Insert(element int) bool {
	if t.Root == nil {
		t.Root = &node{
			Value: element,
			Left:  nil,
			Right: nil,
		}
		t.size++
		return true
	}

	curr := t.Root
	prev := curr
	for curr != nil {
		prev = curr
		if element > curr.Value {
			curr = curr.Right
		} else if element < curr.Value {
			curr = curr.Left
		} else {
			return false
		}
	}

	node := &node{
		Value: element,
		Left:  nil,
		Right: nil,
	}

	if element > prev.Value {
		prev.Right = node
	} else {
		prev.Left = node
	}
	t.size++
	return true
}

// Check if element with value = key exists
func (t *BinarySearchTree) Contains(key int) bool {
	curr := t.Root
	for curr != nil {
		if key == curr.Value {
			return true
		} else if key < curr.Value {
			curr = curr.Left
		} else if key > curr.Value {
			curr = curr.Right
		}
	}
	return false
}

// Delete element with value = key returns false if key does not exists
func (t *BinarySearchTree) Delete(key int) bool {
	if t.Root == nil {
		return false
	}

	curr := t.Root
	prev := curr
	for curr != nil {
		if key > curr.Value {
			prev = curr
			curr = curr.Right
		} else if key < curr.Value {
			prev = curr
			curr = curr.Left
		} else {
			break
		}
	}

	if curr == nil {
		return false
	}

	// For deletion there is various cases
	// case 1: node is a leaf
	if curr.Left == nil && curr.Right == nil {
		if prev.Left == curr {
			prev.Left = nil
		} else if prev.Right == curr {
			prev.Right = nil
		}
		return true
	}
	// case 2: node only has one child
	if curr.Left == nil || curr.Right == nil {
		if prev.Left == curr {
			if curr.Left != nil {
				prev.Left = curr.Left
			} else {
				prev.Left = curr.Right
			}
		} else if prev.Right == curr {
			if curr.Left != nil {
				prev.Right = curr.Left
			} else {
				prev.Right = curr.Right
			}
		}
		return true
	}
	// case 3: worst case node has 2 children in this case find the biggest node in the
	// left subtree and replace the node to be deleted with that node
	prev2 := curr.Left
	curr2 := curr.Left

	for curr2.Right != nil {
		prev2 = curr2
		curr2 = curr2.Right
	}

	if curr2.Left != nil {
		prev2.Right = curr2.Left
	} else {
		prev2.Right = nil
	}

	curr.Value = curr2.Value
	return true

}

// Iterartor returns an iterator that allows to traverse the treee
// @params order-  specifies order of traversal. [ in, pre, post]
// Defaults to in-oder traversal
func (t *BinarySearchTree) Iterator(order string) Iterator {
	switch order {
	case "in":
		it := &InOrderIterartor{
			s:       make([]*node, 0),
			visited: make(map[int]*node),
		}

		if t.Root != nil {
			it.s = append(it.s, t.Root)
		}

		return it
	case "pre":
		it := &PreOrderIterartor{
			s: make([]*node, 0),
		}

		if t.Root != nil {
			it.s = append(it.s, t.Root)
		}

		return it
	case "post":
		it := &PostOrderIterartor{
			s:       make([]*node, 0),
			visited: make(map[int]*node),
		}

		if t.Root != nil {
			it.s = append(it.s, t.Root)
		}

		return it
	default:
		it := &InOrderIterartor{
			s: make([]*node, 0),
		}

		if t.Root != nil {
			it.s = append(it.s, t.Root)
		}

		return it
	}
}
