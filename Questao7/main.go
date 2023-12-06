package main

import "fmt"

// 7. Escreva uma versão do INSERTIONSORT que trabalhe sobre listas encadeadas.
// Ou seja, o mesmo comportamento do algoritmo para vetores deve ser reproduzido pelo seu algoritmo, porém manipulando
// nós de uma lista. Se você considerar que o encadeamento duplo melhora a eficiência do seu algoritmo,
// considere a lista de entrada como duplamente encadeada.

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func inserir(L **Node, x int) {
	novoNó := &Node{Data: x}

	if *L == nil {
		*L = novoNó
		return
	}

	nó := *L
	for nó.Next != nil {
		nó = nó.Next
	}

	nó.Next = novoNó
	novoNó.Prev = nó
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func insertionSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var listaOrdenada *Node
	current := head

	for current != nil {
		next := current.Next

		listaOrdenada = insertSorted(listaOrdenada, current)

		current = next
	}

	return listaOrdenada
}

func insertSorted(listaOrdenada *Node, novoNó *Node) *Node {
	if listaOrdenada == nil || novoNó.Data < listaOrdenada.Data {
		novoNó.Next = listaOrdenada
		if listaOrdenada != nil {
			listaOrdenada.Prev = novoNó
		}
		return novoNó
	}

	current := listaOrdenada
	for current.Next != nil && current.Next.Data < novoNó.Data {
		current = current.Next
	}

	novoNó.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = novoNó
	}

	current.Next = novoNó
	novoNó.Prev = current

	return listaOrdenada
}

func main() {
	var lista *Node

	inserir(&lista, 3)
	inserir(&lista, 1)
	inserir(&lista, 2)
	inserir(&lista, 2)
	inserir(&lista, 5)

	fmt.Println("Lista original:")
	printList(lista)

	listaOrdenada := insertionSort(lista)

	fmt.Println("Lista ordenada:")
	printList(listaOrdenada)
}
