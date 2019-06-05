package main

import (
	"fmt"
	// "math"
)

type Node struct {
	L *Node
	R *Node
	v byte
}

var root *Node
var input = []byte("cdba")

func Insert(n *Node, data byte) *Node {
	if n == nil {
		n = &Node{nil, nil, data}
		return n
	}
	if data < n.v {
		n.L = Insert(n.L, data)
		return n
	}
	n.R = Insert(n.R, data)
	return n
}

func Search(n *Node, data byte) *Node {
	if n == nil {
		return nil
	} else if n.v == data {
		return n
	} else if data < n.v {
		return Search(n.L, data)
	} else {
		return Search(n.R, data)
	}
}

func Min(n *Node) *Node {
	if n == nil || n.L == nil {
		return n
	}
	return Min(n.L)
}

func Max(n *Node) *Node {
	if n == nil || n.R == nil {
		return n
	}
	return Max(n.R)
}

func Print(n *Node, level int) {
    if n != nil {
        Print(n.R, level+1)
        for i := 0; i < level; i++ {
        	fmt.Print("		")
        }
        fmt.Printf("---[ %d\n", n.v)
        Print(n.L, level+1)
    }
}

func Delete(n *Node, data byte) *Node {
	if n == nil {
		return nil
	}
	if data < n.v {
		n.L = Delete(n.L, data)
		return n
	}
	if data > n.v {
		n.R = Delete(n.R, data)
		return n
	}
	if n.L == nil && n.R == nil {
		n = nil
		return nil
	}
	if n.R == nil {
		n = n.L
		return n
	}
	if n.L == nil {
		n = n.R
		return n
	}
	// find smallest value on right side of n
	leftmostR := n.R
	for {
		if leftmostR != nil && leftmostR.L != nil {
			leftmostR = leftmostR.L
		} else {
			break
		}
	}
	n.v = leftmostR.v
	n.R = Delete(n.R, n.v)
	return n
}




func main() {
	fmt.Println(input)
	for i := 0; i < len(input); i++ {
		root = Insert(root, input[i])
	}
	fmt.Println(root)
	Print(root, 0)
	fmt.Println(Search(root, 'X'))
	fmt.Println(Min(root))
	fmt.Println(Max(root))
	Delete(root, 'a')
	Print(root, 0)
}


