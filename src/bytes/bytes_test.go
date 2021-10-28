/*
@Time : 2021/10/25 14:35
@Author : wyf
@File : bytes_test
@Software: GoLand
*/

package bytes

import (
	"fmt"
	"testing"
)



func TestBytesBufferWriteString(t *testing.T) {
	strs := []string{
		"safly",
		"xinmu",
	}
	ret := BytesBufferWriteString(strs...)
	fmt.Println(ret)
}


func TestBytesBufferReadFrom(t *testing.T) {
	retStr := BytesBufferReadFrom()
	fmt.Println(retStr)
}

func TestBytesBufferWriteTo(t *testing.T) {
	BytesBufferWriteTo("write to 输出")
}

func TestBytesBufferRead(t *testing.T)  {
	BytesBufferRead()
}


func TestBytesBufferReadByte(t *testing.T)  {
	BytesBufferReadByte()
}

func TestBytesBufferReadBytes(t *testing.T)  {
	BytesBufferReadBytes()
}


func TestBytesBufferReadString(t *testing.T)  {
	BytesBufferReadString()
}


func TestBytesBufferNext(t *testing.T)  {
	BytesBufferNext()
}


//read

func TestByteReader(t *testing.T)  {
	ByteReader()
}
