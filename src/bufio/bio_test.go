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

	fmt.Println("log...",ret)
}

func TestWriteToBuffer(t *testing.T) {

	bys1 := []byte("hello go")
	bys2 := []byte("hello go")

	ret, _ := WriteToBuffer([][]byte{bys1, bys2}...)

	fmt.Println("log...",ret)
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
