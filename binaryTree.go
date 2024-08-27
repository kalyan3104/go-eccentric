package main

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

// AVLTree represents an AVL tree
type AVLTree struct {
	root *Node
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// height returns the height of the node
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// rightRotate performs a right rotation on the subtree rooted with y
func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	// Return new root
	return x
}

// leftRotate performs a left rotation on the subtree rooted with x
func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	// Return new root
	return y
}

// getBalanceFactor returns the balance factor of the node
func getBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

// insertNode inserts a new node with the given value into the subtree rooted with node
// and returns the new root of the subtree
func insertNode(node *Node, value int) *Node {
	// Perform the normal BST insertion
	if node == nil {
		return &Node{value: value, height: 1}
	}
	if value < node.value {
		node.left = insertNode(node.left, value)
	} else if value > node.value {
		node.right = insertNode(node.right, value)
	} else { // Duplicate values are not allowed in the BST
		return node
	}

	// Update the height of this ancestor node
	node.height = 1 + max(height(node.left), height(node.right))

	// Get the balance factor of this ancestor node to check whether
	// this node became unbalanced
	balance := getBalanceFactor(node)

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if balance > 1 && value < node.left.value {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && value > node.right.value {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	// Return the (unchanged) node pointer
	return node
}

// Insert inserts a new node with the given value into the AVL tree
func (tree *AVLTree) Insert(value int) {
	tree.root = insertNode(tree.root, value)
}

// inOrderTraversal traverses the tree in in-order fashion
func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Printf("%d ", node.value)
		inOrderTraversal(node.right)
	}
}

// preOrderTraversal traverses the tree in pre-order fashion
func preOrderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrderTraversal(node.left)
		preOrderTraversal(node.right)
	}
}

// postOrderTraversal traverses the tree in post-order fashion
func postOrderTraversal(node *Node) {
	if node != nil {
		postOrderTraversal(node.left)
		postOrderTraversal(node.right)
		fmt.Printf("%d ", node.value)
	}
}

// Search searches for a node with a given value in the AVL tree
func (tree *AVLTree) Search(value int) bool {
	return searchNode(tree.root, value)
}

// searchNode searches for a node with a given value in the subtree rooted with node
func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value < node.value {
		return searchNode(node.left, value)
	} else if value > node.value {
		return searchNode(node.right, value)
	}
	return true
}

// minValueNode returns the node with the minimum value found in the tree
func minValueNode(node *Node) *Node {
	current := node

	// Loop down to find the leftmost leaf
	for current.left != nil {
		current = current.left
	}
	return current
}

// deleteNode deletes a node with the given value from the subtree rooted with node
// and returns the new root of the subtree
func deleteNode(root *Node, value int) *Node {
	// Perform standard BST delete
	if root == nil {
		return root
	}

	if value < root.value {
		root.left = deleteNode(root.left, value)
	} else if value > root.value {
		root.right = deleteNode(root.right, value)
	} else {
		// Node with only one child or no child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		// Node with two children: Get the inorder successor (smallest in the right subtree)
		temp := minValueNode(root.right)

		// Copy the inorder successor's content to this node
		root.value = temp.value

		// Delete the inorder successor
		root.right = deleteNode(root.right, temp.value)
	}

	// Update the height of the current node
	root.height = 1 + max(height(root.left), height(root.right))

	// Get the balance factor of this node
	balance := getBalanceFactor(root)

	// Balance the tree

	// Left Left Case
	if balance > 1 && getBalanceFactor(root.left) >= 0 {
		return rightRotate(root)
	}

	// Left Right Case
	if balance > 1 && getBalanceFactor(root.left) < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	// Right Right Case
	if balance < -1 && getBalanceFactor(root.right) <= 0 {
		return leftRotate(root)
	}

	// Right Left Case
	if balance < -1 && getBalanceFactor(root.right) > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

// Delete deletes a node with the given value from the AVL tree
func (tree *AVLTree) Delete(value int) {
	tree.root = deleteNode(tree.root, value)
}

func main() {
	tree := &AVLTree{}

	values := []int{10, 20, 30, 40, 50, 25}
	for _, value := range values {
		tree.Insert(value)
	}

	fmt.Println("In-order traversal of the AVL tree:")
	inOrderTraversal(tree.root)
	fmt.Println()

	fmt.Println("Pre-order traversal of the AVL tree:")
	preOrderTraversal(tree.root)
	fmt.Println()

	fmt.Println("Post-order traversal of the AVL tree:")
	postOrderTraversal(tree.root)
	fmt.Println()

	fmt.Printf("Searching for value 25 in the AVL tree: %v\n", tree.Search(25))
	fmt.Printf("Searching for value 60 in the AVL tree: %v\n", tree.Search(60))

	fmt.Println("Deleting value 30 from the AVL tree")
	tree.Delete(30)

	fmt.Println("In-order traversal after deletion:")
	inOrderTraversal(tree.root)
	fmt.Println()
}
