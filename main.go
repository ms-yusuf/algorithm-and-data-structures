package main

import (
	"github.com/ms-yusuf/algorithm-and-data-structures/datastructures"
	"fmt"
)

func main(){
	ll := datastructures.LinkedList{}
	ll.AddToFront("Test1")
	ll.AddToFront("Test2")
	ll.AddToFront("head")
	ll.AddAfter("Test1.5", "Test2")
	ll.Iterate()
	ll.AddToEnd("End")
	fmt.Println("\n\nFind 1st occur:", ll.FindFirstOccurrence("Test2").Data)
	lastNode := ll.GetLastNode()
	firstNode := ll.GetFirstNode()
	fmt.Println(lastNode.Data)
	fmt.Println(firstNode.Data)
}
