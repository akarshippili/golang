package datastructures

import (
	"fmt"
	"unsafe"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BtreeTest() {

	left := TreeNode{
		Val: 10,
	}

	right := TreeNode{
		Val: 10,
	}

	temp := TreeNode{
		Val: -1000,
	}

	node := TreeNode{
		Val:   1,
		Left:  &left,
		Right: &right,
	}

	ptrnode := &node
	fmt.Println(node, node.Left, node.Right)
	fmt.Println(unsafe.Sizeof(node), unsafe.Sizeof(ptrnode))

	node.Left = &temp
	fmt.Println(node, node.Left, node.Right)

	*node.Right = temp
	fmt.Println(node, node.Left, node.Right)

	// a := []float64{1}
	// ptra := &a
	// fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(ptra))
	// fmt.Println(reflect.TypeOf(a), reflect.TypeOf(ptra))

	// course := models.Course{
	// 	Author: "a",
	// 	Level:  "b",
	// 	Rating: 1.11,
	// }
	// ptrcourse := &course
	// fmt.Println(unsafe.Sizeof(course), unsafe.Sizeof(ptrcourse))
	// fmt.Println(reflect.TypeOf(course), reflect.TypeOf(ptrcourse))

}
