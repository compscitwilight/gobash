package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	parent   *TreeNode
	children []TreeNode
	content  string
}

func create_tree() TreeNode {
	root_node := TreeNode{}
	return root_node
}

func append_child_node(t TreeNode, child *TreeNode) TreeNode {
	child_node := *child
	child_node.parent = &t
	t.children = append(t.children, child_node)
	//t.children[len(t.children)] = child_node
	return t
}

func log_tree(t TreeNode) {
	iterate := func(n TreeNode) {
		node := n
		ready := false
		log.Println(n.content)
		for _, child := range node.children {
			ready = false
			log.Println(fmt.Sprintf("├ %v", child.content))
			if len(child.children) > 0 && !ready {
				node = child
				ready = true
			}
		}
	}

	iterate(t)
}
