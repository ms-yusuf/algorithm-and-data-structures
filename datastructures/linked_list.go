package datastructures 

import "fmt"

type Node struct {
	Data string
	NextNode *Node
}

type LinkedList struct {
	FirstNode *Node
}

func (l *LinkedList) AddToFront(d string){
	Node := Node{Data: d}
	if Node.NextNode != nil {
		Node.NextNode = l.FirstNode
	}
	if l.FirstNode != &Node {
		Node.NextNode = l.FirstNode
	}
	l.FirstNode = &Node
}

func (l *LinkedList) AddToEnd(d string){
	node := Node{Data: d, NextNode: nil}
	lastNode := l.GetLastNode()
	lastNode.NextNode = &node
}

func (l *LinkedList) AddAfter(d, f string) {
	foundNode := l.FindFirstOccurrence(f)
	if foundNode != nil {
		n := Node{d, foundNode.NextNode}
		foundNode.NextNode = &n
	}
}


func (l *LinkedList) FindFirstOccurrence(f string) *Node {
	for n := l.FirstNode; n != nil; n = n.NextNode {
		if n.Data == f {
			return n
		}
	}
	return nil
}

func (l *LinkedList) Iterate(){
	for n := l.FirstNode; n != nil; n = n.NextNode {
		fmt.Println(n.Data);
	}
}

func (l *LinkedList) GetFirstNode() *Node {
	return l.FirstNode
}

func (l *LinkedList) GetLastNode() *Node {
	var lastNode *Node
	
	for n := l.FirstNode; n != nil; n = n.NextNode {
		if n.NextNode == nil {
			lastNode = n
		}
	}
	return lastNode
}
