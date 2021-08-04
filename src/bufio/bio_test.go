/*
@Time : 2020-12-18 15:12
@Author : wyf
@File : bio_test
@Software: GoLand
*/

package bufio

import (
	"fmt"
	"testing"
)

//读取数据到buffer中
func TestReadBufferToBuf(t *testing.T) {

	bys := []byte("hello go")

	ret, _ := ReadBufferToBuf(bys)

	fmt.Println(ret)
}

func TestWriteToBuffer(t *testing.T) {

	bys1 := []byte("hello go")
	bys2 := []byte("hello go")

	ret, _ := WriteToBuffer([][]byte{bys1, bys2}...)

	fmt.Println(ret)
}
