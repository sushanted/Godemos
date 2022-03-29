package main

import "fmt"

type Value interface {
	String() string
	Add(value Value) Value
}

// Specific implementations

type IntValue int

func (intValue IntValue) String() string {
	return fmt.Sprintf("%d", int(intValue))
}

func (intValue IntValue) Add(value Value) Value {
	return intValue + value.(IntValue)
}

type StringValue string

func (stringValue StringValue) String() string {
	return string(stringValue)
}

func (stringValue StringValue) Add(value Value) Value {
	return stringValue + value.(StringValue)
}

type List []int

func (list List) String() string {
	return fmt.Sprintf("%v", []int(list))
}

func (list List) Add(value Value) Value {

	otherList := value.(List)

	smallList := list
	largeList := otherList

	if len(largeList) < len(smallList) {
		smallList, largeList = largeList, smallList
	}

	addedList := List(make([]int, len(largeList)))
	copy(addedList, largeList)

	for i := 0; i < len(smallList); i++ {
		addedList[i] += smallList[i]
	}

	return addedList
}

// generic data structure and algorithm

type Node struct {
	Value
	Left  *Node
	Right *Node
}

func (node *Node) String() string {

	if node == nil {
		return "_"
	}
	return node.Value.String() + "(" + node.Left.String() + "," + node.Right.String() + ")"

}

func (node *Node) Sum() Value {

	var finalValue Value = node.Value

	if node.Left != nil {
		finalValue = finalValue.Add(node.Left.Sum())
	}

	if node.Right != nil {
		finalValue = finalValue.Add(node.Right.Sum())
	}

	return finalValue
}

func main() {

	intRoot := &Node{IntValue(4), &Node{Value: IntValue(1)}, &Node{Value: IntValue(2)}}
	fmt.Println(intRoot.String())
	fmt.Println(intRoot.Sum())

	stringRoot := &Node{StringValue("A"), &Node{Value: StringValue("B")}, &Node{Value: StringValue("C"), Left: &Node{Value: StringValue("D")}}}
	fmt.Println(stringRoot.String())
	fmt.Println(stringRoot.Sum())

	listRoot := &Node{List([]int{1, 2, 3, 4}), &Node{Value: List([]int{3, 4})}, &Node{Value: List([]int{7, 8, 9})}}
	fmt.Println(listRoot.String())
	fmt.Println(listRoot.Sum())

}

/**

OUTPUT:
4(1(_,_),2(_,_))
7
A(B(_,_),C(D(_,_),_))
ABCD
[1 2 3 4]([3 4](_,_),[7 8 9](_,_))


*/
