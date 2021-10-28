/*
@Time : 2021/10/25 14:35
@Author : wyf
@File : buffer
@Software: GoLand
*/

package bytes

import (
	"bytes"
	"fmt"
	"os"
)

/***********************buffer*********************************
buffer
 */
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
讲字节流写入到缓存
 */
func BytesBufferWriteString(strs ...string) string {
	buffer := bytes.NewBuffer([]byte{})
	for _,str := range strs{
		//或者
		//buffer.WriteString(str)
		buffer.Write([]byte(str))
	}
	return buffer.String()
}

/**
从文件中读取内容
 */
func BytesBufferReadFrom() string {
	file,err := os.Open("buffer.txt")
	if err!=nil{
		fmt.Println(err)
	}
	buffer := bytes.NewBuffer([]byte{})
	buffer.ReadFrom(file)
	retStr := buffer.String()
	return  retStr
}

/**
写出到文件中
 */
func BytesBufferWriteTo(str string)  {
	file, err := os.OpenFile("write_to.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err!=nil{
		fmt.Println(err)
	}
	buf := bytes.NewBufferString(str)
	buf.WriteTo(file)
}

/**
读到字节数组里面
 */
func BytesBufferRead()  {
	bufRead := bytes.NewBufferString("hello")
	fmt.Println(bufRead.String())

	var sRead = make([]byte, 3)
	bufRead.Read(sRead) //读3个字节到sRead数组里面

	fmt.Println(bufRead.String()) // 因为前三个被读出了,打印结果为lo
	fmt.Println(string(sRead))// 打印结果为hel,因为前三个被读出了
}

/**
ReadByte一个字节一个字节读
 */

func BytesBufferReadByte()  {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	// buf.String()方法是吧buf里的内容转成string，>以便于打印
	b, _ := buf.ReadByte()    // 读取第一个byte，赋值给b
	fmt.Println(buf.String()) // 打印 ello，缓冲器头部第一个h被拿掉
	fmt.Println(string(b))//输出h
}


/**
ReadBytes
读到分隔符返回给变量
*/

func BytesBufferReadBytes()  {
	var delim byte = 'f'
	buf := bytes.NewBufferString("hsaflyllo")
	fmt.Println(buf.String())
	// buf.String()方法是吧buf里的内容转成string，>以便于打印
	b, _ := buf.ReadBytes(delim)
	fmt.Println(buf.String()) // 打印 lyllo
	fmt.Println(string(b))//输出hsaf
}

/**
Readstring
读到分隔符返回给变量
*/

func BytesBufferReadString()  {
	var delim byte = 'f'
	buf := bytes.NewBufferString("hsaflyllo")
	fmt.Println(buf.String())
	// buf.String()方法是吧buf里的内容转成string，>以便于打印
	b, _ := buf.ReadString(delim)
	fmt.Println(buf.String()) // 打印 lyllo
	fmt.Println(b)//输出hsaf
}

/**Next
 */
func BytesBufferNext()  {
	buf := bytes.NewBufferString("hsaflyllo")
	b := buf.Next(3)
	fmt.Println(buf.String())
	fmt.Println(string(b))
}


/********************************************************
bytes
*/




/************************reader********************************
reader 读取字节数组到字节数组，利用了copy功能
*/

func ByteReader()  {
	b1 := []byte("Hello World!")
	rd := bytes.NewReader(b1)

	byteRet1,_ := rd.ReadByte()
	fmt.Println(string(byteRet1))
	byteRet2,_ := rd.ReadByte()
	fmt.Println(string(byteRet2))

	readBytes :=make([]byte, 6)
	rd.Read(readBytes)
	fmt.Println(string(readBytes))

}
