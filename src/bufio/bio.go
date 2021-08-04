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
