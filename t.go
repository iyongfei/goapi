/*
@Time : 2020-12-17 00:29
@Author : wyf
@File : t
@Software: GoLand
*/

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)


var HexMap = map[string]string{
	"20":" ",
	"21":"!",
	"22":"\"",
	"23":"#",
	"24":"$",
	"25":"%",
	"26":"&",
	"27":"'",
	"28":"(",
	"29":")",

	"2A":"*",
	"2B":"+",
	"2C":",",
	"2D":"-",
	"2E":".",
	"2F":"/",
	"30":"0",
	"31":"1",
	"32":"2",
	"33":"3",
	"34":"4",
	"35":"5",
	"36":"6",
	"37":"7",
	"38":"8",
	"39":"9",
	"3A":":",
	"3B":";",
	"3C":"<",
	"3D":"=",
	"3E":">",
	"3F":"?",
	"40":"@",
	"41":"A",
	"42":"B",
	"43":"C",
	"44":"D",
	"45":"E",
	"46":"F",
	"47":"G",
	"48":"H",
	"49":"I",
	"4A":"J",
	"4B":"K",
	"4C":"L",
	"4D":"M",
	"4E":"N",
	"4F":"O",
	"50":"P",
	"51":"Q",
	"52":"R",
	"53":"S",
	"54":"T",
	"55":"U",
	"56":"V",
	"57":"W",
	"58":"X",
	"59":"Y",
	"5A":"Z",
	"5B":"[",
	"5C":"\\",
	"5D":"]",
	"5E":"^",
	"5F":"_",
	"60":"`",
	"61":"a",
	"62":"b",
	"63":"c",
	"64":"d",
	"65":"e",
	"66":"f",
	"67":"g",
	"68":"h",
	"69":"i",

	"6A":"j",
	"6B":"k",
	"6C":"l",
	"6D":"m",
	"6E":"n",
	"6F":"o",

	"70":"p",
	"71":"q",
	"72":"r",
	"73":"s",
	"74":"t",
	"75":"u",
	"76":"v",
	"77":"w",
	"78":"x",
	"79":"y",

	"7A":"z",
	"7B":"{",
	"7C":"|",
	"7D":"}",
	"7E":"~",
	//"7F":"DEL",
}

//”   ;   :   \
func main() {

	ret:= `Bay " Tech; Au:thentic\ation Bypass`
	fmt.Println(ret)

	s4 := strings.Replace(ret,`\`,`\\`,-1)
	fmt.Println(s4)

	s1 := strings.Replace(s4,`"`,`\"`,-1)
	fmt.Println(s1)


	s2 := strings.Replace(s1,`;`,`\;`,-1)
	fmt.Println(s2)

	s3 := strings.Replace(s2,`:`,`\:`,-1)
	fmt.Println(s3)





}

/**


转换规则：
把16进制字符串，转换为字符，如果不属于ASCII 打印字符，就无需转换利用||包裹起来
例如
607F8F61618F
转换结果如下
`[7F8F]aa[8F]
 */

func M(r string) string {
	tmp := ""
	var buffer bytes.Buffer

	hexTmp := ""
	for k,v:=range r{
		index := k%2
		if index == 1{
			buffer.WriteString(string(v))
			ret := buffer.String()
			if _,ok:=HexMap[ret];ok{
				hex_data, _ := hex.DecodeString(ret)
				retStr:= string(hex_data)
				if hexTmp != ""{
					hexTmp = ""
					//tmp = strings.TrimRight(tmp," ")
					tmp += "]]"
				}
				tmp += retStr
			}else{
				if hexTmp == ""{
					join := "[[" + ret + " "
					hexTmp = join
					tmp += join
					if k== len(r)-1{
						tmp = strings.TrimRight(tmp," ")
						tmp += "]]"
					}
				}else{
					hexTmp +=  ret
					tmp +=  ret + " "
					if k== len(r)-1{
						tmp = strings.TrimRight(tmp," ")
						tmp += "]]"
					}
				}
			}
		}else {
			buffer.Reset()
			buffer.WriteString(string(v))
		}
	}
	return tmp
}
/**
607F8F61618F
转换结果如下
`[7F8F]aa[8F]
 */
