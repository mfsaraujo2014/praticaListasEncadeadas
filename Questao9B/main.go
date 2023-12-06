package main

import "fmt"

// 9-(b) Desta vez sem contar o número de nós da lista,  encontrar o nó que reside no meio da lista.

type Node struct {
	Data int
	Next *Node
}

func encontrarMeioSemContarTotal(head *Node) *Node {
	rapido := head
	lento := head

	for rapido != nil && rapido.Next != nil {
		rapido = rapido.Next.Next
		lento = lento.Next
	}

	return lento
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
	lista := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: nil}}}}

	fmt.Println("Lista original:")
	printList(lista)

	meioSemContarTotal := encontrarMeioSemContarTotal(lista)
	fmt.Printf("Nó do meio sem contar total: %d\n", meioSemContarTotal.Data)
}
