/*
@Time : 2021/12/8 10:32
@Author : wyf
@File : list
@Software: GoLand
*/

package main

import (
	"container/list"
	"fmt"
)

func Example() {
	l := list.New()
	//插入
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	//遍历
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//len
	fmt.Println("len--->",l.Len())

	//move
	l.Remove(e4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//out put
	//1
	//2
	//3
	//4
	//len---> 4
	//1
	//2
	//3

}


func main()  {
	Example()
}
