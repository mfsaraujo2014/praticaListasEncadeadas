package main

import "fmt"

//9-(a) Dada uma lista encadeada 𝐿, encontrar o nó que reside no meio da lista.

type Node struct {
	Data int
	Next *Node
}

func encontrarMeioComContagemTotal(head *Node) *Node {
	totalNodos := 0
	current := head
	for current != nil {
		totalNodos++
		current = current.Next
	}

	meioIndex := totalNodos / 2
	current = head
	for i := 0; i < meioIndex; i++ {
		current = current.Next
	}

	return current
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	lista := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5, Next: nil}}}}}

	fmt.Println("Lista original:")
	printList(lista)

	meioComContagemTotal := encontrarMeioComContagemTotal(lista)
	fmt.Printf("Nó do meio com contagem total: %d\n", meioComContagemTotal.Data)
}
