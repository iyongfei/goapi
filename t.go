/*
@Time : 2020-12-17 00:29
@Author : wyf
@File : t
@Software: GoLand
*/

package main

import "fmt"

type Format int

const (
	_ Format = (iota + 1) * 2
	A
)

func main() {
	fmt.Println(A)

	var buf [3]byte
	fmt.Println("ReadBufferToBuf", len(buf))
	for k, v := range buf {
		fmt.Println(k, v)
	}
}
