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

/*******************write****************
write
*/
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


/*******************reader****************
推荐使用ReadBytes('\n') or ReadString('\n')来替代ReadLine。

readbytes
*/
func BufioReadBytes(str string,split byte) string {
	buffer := strings.NewReader(str)
	reader := bufio.NewReader(buffer)
	token, err := reader.ReadBytes(split)
	if err != nil {
		panic(err)
	}
	return string(token)
}


/**
readString
*/
func BufioReadString(str string,split byte) string {
	buffer := strings.NewReader(str)
	reader := bufio.NewReader(buffer)
	token, err := reader.ReadString(split)
	if err != nil {
		panic(err)
	}
	return token
}




//从buffer中读数据到buffer中
func ReadBufferToBuf(bys []byte) (string, error) {
	buffer := bytes.NewBuffer(bys)
	reader := bufio.NewReader(buffer)
	var buf [128]byte
	lens, err := reader.Read(buf[:])

	return string(buf[:lens]), err
}

func Peek(bys []byte)  {
	buffer := bytes.NewBuffer(bys)
	reader := bufio.NewReader(buffer)
	ret,_ := reader.Peek(3)
	fmt.Println("ret::",string(ret))
}



