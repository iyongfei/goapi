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



/******************ScanScan****************
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
