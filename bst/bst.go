package bst

type node struct {
	left  *node
	right *node
	key   int
	value string
}

type Bst struct {
	root *node
}

func insert(treeNode *node, newNode *node) {
	if newNode.key < treeNode.key {
		if treeNode.left == nil {
			treeNode.left = newNode
		} else {
			insert(treeNode.left, newNode)
		}
	} else {
		if treeNode.right == nil {
			treeNode.right = newNode
		} else {
			insert(treeNode.right, newNode)
		}
	}
}

func (b *Bst) Insert(key int, value string) {
	newNode := &node{nil, nil, key, value}

	if b.root == nil {
		b.root = newNode
		return
	}

	insert(b.root, newNode)
}
