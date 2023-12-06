package main

import "fmt"

// 4. Dados dois nós n1
// e n2 de listas encadeadas, determinar se ambos fazem parte de uma mesma lista.

type Node struct {
	Data int
	Next *Node
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func areNodesInSameList(head *Node, n1 *Node, n2 *Node) bool {
	current := head

	for current != nil {
		if current == n1 || current == n2 {
			return true
		}
		current = current.Next
	}

	return false
}

func main() {
	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5}}}}}

	node1 := head.Next.Next           // Nó com valor 3
	node2 := head.Next.Next.Next.Next // Nó com valor 5

	fmt.Println("Lista original:")
	printList(head)

	if areNodesInSameList(head, node1, node2) {
		fmt.Println("Os nós estão na mesma lista.")
	} else {
		fmt.Println("Os nós não estão na mesma lista.")
	}
}
