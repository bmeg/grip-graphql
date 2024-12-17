package graph

import (
	"fmt"
	"strings"
)

type Node struct {
	value    string
	children map[string]*Node
	fields   []string
	visited  bool
}

func createNodeGraph(paths []string) *Node {
	/* Construct a Node graph from a list of '.' delimited traversal paths */
	graph := &Node{
		children: make(map[string]*Node),
	}

	for _, path := range paths {
		parts := strings.Split(path, ".")
		current := graph

		for i, part := range parts {
			if strings.HasSuffix(part, "Type") && !strings.HasSuffix(part, "resourceType") {
				if _, ok := current.children[part]; !ok {
					current.children[part] = &Node{
						value:    part,
						children: make(map[string]*Node),
					}
				}
				current = current.children[part]
			} else {
				current.fields = append(current.fields, strings.Join(parts[i:], "."))
				break
			}
		}
	}
	return graph
}

func constructTypeTraversal(paths []string) (string, map[string][]string) {
	//Build traversal using a modifed Depth First Search algorithm
	typeFields := make(map[string][]string)
	returnPath, traversalPath := []string{}, []string{}
	stack := []*Node{}

	graph := createNodeGraph(paths)

	for node := range graph.children {
		if strings.HasSuffix(node, "Type") {
			stack = append(stack, graph.children[node])
		}
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.visited {
			continue
		}
		current.visited = true

		// Append fields on each node to typeFields
		if len(current.fields) > 0 {
			typeFields[current.value] = append(typeFields[current.value], current.fields...)
		}

		// Locate parent node and dertermine if traversal path needs to backtrack or not
		parentValue := findParentNode(graph, current)
		if parentValue != "" && (len(traversalPath) > 0 && traversalPath[len(traversalPath)-1] != parentValue) {
			traversalPath = append(traversalPath, parentValue)
			returnPath = append(returnPath, "SELECT_"+parentValue)
		}

		traversalPath = append(traversalPath, current.value)
		if len(traversalPath) > 1 {
			returnPath = append(returnPath, "OUTNULL_"+current.value)
		}

		for _, child := range current.children {
			stack = append(stack, child)
		}
	}

	return strings.Join(returnPath, "."), typeFields
}

func findParentNode(graph *Node, node *Node) string {
	// Traverse the entire graph to find a node that has the given node as a child
	var findParent func(current *Node) string
	findParent = func(current *Node) string {
		if current == nil {
			return ""
		}
		for _, child := range current.children {
			if child == node {
				return current.value
			}
			parentValue := findParent(child)
			if parentValue != "" {
				return parentValue
			}
		}
		return ""
	}
	return findParent(graph)
}

func printNode(node *Node, depth int) {
	/* Utility funciton used for printing the contents
	   of the node struct in a yaml style format */
	if node == nil {
		return
	}
	fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.value)
	for _, child := range node.children {
		printNode(child, depth+1)
	}
}
