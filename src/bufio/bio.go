/*
@Time : 2020-12-18 14:53
@Author : wyf
@File : bio
@Software: GoLand
*/

package bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//从buffer中读数据到buffer中
func ReadBufferToBuf(bys []byte) (string, error) {
	buffer := bytes.NewBuffer(bys)
	reader := bufio.NewReader(buffer)
	var buf [128]byte
	lens, err := reader.Read(buf[:])

	return string(buf[:lens]), err
}

//读数据到writer buffer中
func WriteToBuffer(bys ...[]byte) (string, error) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	var err error
	for _, v := range bys {
		_, err = writer.Write(v)
	}
	err = writer.Flush()
	return string(buffer.Bytes()), err
}

/**
peek
 */

func Peek(bys []byte)  {
	buffer := bytes.NewBuffer(bys)
	reader := bufio.NewReader(buffer)
	ret,_ := reader.Peek(3)
	fmt.Println("ret::",string(ret))
}
/**
扫描\n
 */
func ScannerLine(str string)  {
	scanner:=bufio.NewScanner(
		strings.NewReader(str),
	)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

/**
扫描空格
*/
func ScannerWords(str string)  {
	scanner:=bufio.NewScanner(
		strings.NewReader(str),
	)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



/**
扫描字节数组
*/
func ScannerBytes(str string)  {
	scanner:=bufio.NewScanner(
		strings.NewReader(str),
	)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



func ScannerCustome(str string)  {
	scanner:=bufio.NewScanner(
		strings.NewReader(str),
	)
	scanner.Split(ScanCustomSplitFunc)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


func ScanCustomSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("safly")); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
