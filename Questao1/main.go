package main

import "fmt"

// 1. Escreva um algoritmo de remoção que, dada uma lista encadeada L e um valor de entrada x, remove
// de L todos os valores menores que x

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

func removeElements(head *Node, x int) *Node {
	dummy := &Node{Data: 0, Next: head}
	prev, current := dummy, head

	for current != nil {
		if current.Data < x {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return dummy.Next
}

func main() {

	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5}}}}}

	fmt.Println("Lista original:")
	printList(head)

	head = removeElements(head, 3)

	fmt.Println("Lista após a remoção:")
	printList(head)
}
