/*
@Time : 2021/10/25 14:35
@Author : wyf
@File : buffer
@Software: GoLand
*/

package bytes

import (
	"bytes"
)

/**
buf1 := bytes.NewBufferString("hello")
	fmt.Println(buf1)
	buf2 := bytes.NewBuffer([]byte("hello"))
	fmt.Println(buf2)
	buf3 := bytes.NewBuffer([]byte{byte('h'), byte('e'), byte('l'), byte('l'), byte('o')})
	fmt.Println(buf3)
    // 以上三者等效

	buf4 := bytes.NewBufferString("")
	fmt.Println(buf4)
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf5)
    // 以上两者等效
 */


/**
WriteString
或者
Write
 */
func BytesBufferWriteString(strs ...string) string {
	buffer := bytes.NewBuffer([]byte{})

	for _,str := range strs{
		//buffer.WriteString(str)
		buffer.Write([]byte(str))
	}
	return buffer.String()
}
/**

ReadString
ReadBytes
ReadByte
Next
Read
WriteTo
ReadFrom
Write
 */
