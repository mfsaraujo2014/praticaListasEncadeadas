package main

import "fmt"

// 3. Escreva um algoritmo que, dadas duas listas encadeadas como entrada,
// retorna uma nova lista concatena a segunda lista no fim da primeira.

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

func concatenateLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}

	current := list1
	for current.Next != nil {
		current = current.Next
	}

	current.Next = list2
	return list1
}

func main() {
	lista1 := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3}}}
	lista2 := &Node{Data: 4, Next: &Node{Data: 5}}

	fmt.Println("Lista 1:")
	printList(lista1)

	fmt.Println("Lista 2:")
	printList(lista2)

	listaConcatenada := concatenateLists(lista1, lista2)

	fmt.Println("Lista Concatenada:")
	printList(listaConcatenada)
}
