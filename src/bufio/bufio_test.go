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


/**
readbytes
类似于slice切片
*/
func TestBufioReadBytes(t *testing.T) {
	str := "ABCDHEFGHIJKELM"
	ret := BufioReadBytes(str,byte('H'))
	fmt.Println(ret)
}

/**
readbytes
类似于slice切片
*/
func TestBufioReadString(t *testing.T) {
	str := "ABCDHEFGHIJKELM"
	ret := BufioReadString(str,byte('H'))
	fmt.Println(ret)
}

//读取数据到buffer中
func TestReadBufferToBuf(t *testing.T) {
	bys := []byte("hello go")
	ret, _ := ReadBufferToBuf(bys)
	fmt.Println("log...",ret)
}

func TestWriteToBuffer(t *testing.T) {

	bys1 := []byte("hello go")
	bys2 := []byte("hello go")

	ret, _ := WriteToBuffer([][]byte{bys1, bys2}...)

	fmt.Println("log...",ret)
}

func TestBufioString(t *testing.T) {
	ret:= BufioString("rrrree的")
	fmt.Println("log...",ret)
}




//读取数据到buffer中
func TestBufioPeek(t *testing.T) {
	bys := []byte("hello go")
	Peek(bys)
}


/**
默认按着\n分割
 */

func TestWriteToScannerLine(t *testing.T) {
	ret := "ABCDEFG\nHIJKELM"
	ScannerLine(ret)
}

/**
按空格分割
 */
func TestWriteToScannerByWords(t *testing.T) {
	ret := "ABCDEFGsafl yHIJKELM"
	ScannerWords(ret)
}

/**
按读取每一个字符
*/
func TestWriteToScannerByBytes(t *testing.T) {
	ret := "ABCDEFGsafl yHIJKELM"
	ScannerBytes(ret)
}

/**
自定义切分
 */
func TestWriteToScannerByCustome(t *testing.T) {
	ret := "ABCDEFGsaflyHIJKELM"
	ScannerCustome(ret)
}
