package main

import "fmt"

// 2. Escreva um algoritmo que, dada uma lista simplesmente encadeada, inverte o encadeamento da lista.
// Ou seja, o primeiro elemento passará a ser o último, o segundo passará a ser o penúltimo, etc.

// Definindo a estrutura de um nó na lista encadeada
type Node struct {
	Data int
	Next *Node
}

// Função para imprimir a lista encadeada
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func reverseList(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5}}}}}

	fmt.Println("Lista original:")
	printList(head)

	head = reverseList(head)

	fmt.Println("Lista invertida:")
	printList(head)
}
