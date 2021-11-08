/*
@Time : 2021/10/25 19:03
@Author : wyf
@File : util
@Software: GoLand
*/

package util

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)


func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err!=nil{
		return false
	}
	return s.IsDir()
}


func readf(path string, fileMode int) []byte {
	var file *os.File
	if fileMode == 1 {
		file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	} else if fileMode == 2 {
		file, _ = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	} else {
		file, _ = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	}
	c, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read err:%+v\n", err)
	}
	return c
}

func writef(path string, fileMode int, data []byte) {
	var file *os.File
	if fileMode == 1 {
		file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	} else if fileMode == 2 {
		file, _ = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	} else {
		file, _ = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	}
	file.Write(data)
}


func FileMode(path string, fileMode int) (*os.File,error){
	var err error
	var file *os.File
	if fileMode == 1 {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	} else if fileMode == 2 {
		file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	} else {
		file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	}
	return file,err
}

var HexMap = map[string]string{
	"20":" ",
	"21":"!",
	//"22":"\"",
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
	//"3A":":",
	//"3B":";",
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
	//"7C":"|",
	"7D":"}",
	"7E":"~",
	"7F":"DEL",
}

/**
转换规则：
把16进制字符串，转换为字符，如果不属于ASCII 打印字符，就无需转换利用||包裹起来
例如
607F8F61618F
转换结果如下
`[7F8F]aa[8F]
*/

func Heco(r string) string {
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
					tmp += "||"
				}
				tmp += retStr
			}else{
				if hexTmp == ""{
					join := "||" + ret + " "
					hexTmp = join
					tmp += join
					if k== len(r)-1{
						tmp = strings.TrimRight(tmp," ")
						tmp += "||"
					}
				}else{
					hexTmp +=  ret
					tmp +=  ret + " "
					if k== len(r)-1{
						tmp = strings.TrimRight(tmp," ")
						tmp += "||"
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

var RuleHeaderMap = map[string]string{
	"ftp":   "#ftp:4011000~4011999",
	"telnet":"#telnet:4012000~4012999",
	"tftp":"#tftp:4013000~4013999",
	"imap":"#imap:4014000~4014999",//多一个
	"smtp":"#smtp:4015000~4015999",//多一个
	"dhcp":"#dhcp:4016000~4016999",
	"dnp3":"#dnp3:4017000~4017999",//有误
	"dns":"#dns:4018000~4018999",
	"http":"#http:4019000~4019999",
	"icmp":"#icmp:4020000~4020999",
	"kerberos":"#kerberos:4021000~4021999",
	"netbios":"#netbios:4022000~4022999",
	"nfs":"#nfs:4023000~4023999",
	"ntp":"#ntp:4024000~4024999",//错误
	"smb":"#smb:4025000~4025999",
	"ssl":"#ssl:4026000~4026999",
	"snmp":"#snmp:4027000~4027999",
	"ssh":"#ssh:4028000~4028999",
	"tls":"#tls:4029000~4029999",
	"xss":"#xss:4030000~4030999\t",
	"shellcode":"#shellcode:4031000~4031999",
	"sql":"#sql:4032000~4032999",
	//#activex:4033000~4033999
	"dos":"#dos:4034000~4034999",
	"bypass":"#bypass:4035000~4035999",
	"chat":"#chat:4036000~4036999",
	"cmi":"#cmi:4037000~4037999",
	"flood":"#flood:4038000~4038999",
	"fuzz":"#fuzz:4039000~4039999",
	"game":"#game:4040000~4040999",
	"hunting":"#hunting:4041000~4041999",
	"info":"#info:4042000~4042999",
	"sip":"#sip:4043000~4043999",
	"trojan":"#trojan:4044000~4044999",
	"worm":"#worm:4045000~4045999",
	"portscan":"#portscan:4046000~4046999",
	"iis":"#iis:4050000~4050999",
	"js":"#js:4051000~4051999",
	"noop":"#noop:4052000~4052999",
	"oracle":"#oracle:4053000~4053999",
	"overflow":"#overflow:4054000~4054999",
	"php":"#php:4055000~4055999",
	"policy":"#policy:4056000~4056999",
	"vbs":"#vbs:4057000~4057999",
	"web":"#web:4058000~4058999",
	"microsoft":"#microsoft:4059000~4059999",
	"denial_of_service":"#denial_of_service:4060000~4060999",
	"other":"#other:4061000~4061999",
}

func GetRuleLineHeaderMap() map[string] string{
	f, err := os.Open("suricata_rules_cnt.tex")
	if err != nil {
		return nil
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return nil
		}

	}
	return nil
}
