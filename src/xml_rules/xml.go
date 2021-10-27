/*
@Time : 2021/10/25 18:31
@Author : wyf
@File : xml
@Software: GoLand
*/

package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"goapi/src/xml_rules/util"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	XmlName = "XmlName"
	ExecDesc = "ExecDesc"
	EngDesc = "EngDesc"
	Protocol = "Protocol"
	Payload = "Payload"
)
var (
	intrusionPath = "rules/microsoft"
)

//   worm
//创建目录

func init()  {
	baseXmlDir := "excel"

	if !util.Exists(baseXmlDir){
		_= os.Mkdir(baseXmlDir, os.ModePerm)
	}
}


func main()  {
	var ruleDir string
	flag.StringVar(&ruleDir, "v", "", "vid")
	flag.Parse()

	if ruleDir != "" {
		if !strings.HasPrefix(ruleDir,"rules/"){
			fmt.Printf("input rule:%s right fomate",ruleDir)
			return
		}
		intrusionPath = ruleDir
	}else{
		fmt.Printf("input rule:%s content null",ruleDir)
		return
	}

	//当前的xml，要插入的csv文件路径
	segIndex := strings.Index(intrusionPath,"/")
	ruleDirName:= intrusionPath[segIndex+1:]
	csvPath := "excel" + "/" + ruleDirName + ".csv"

	header := [][]string{
		[]string{"XmlName","ExecDesc","EngDesc","Protocol","SPECProtocol","Payload",},
	}
	insertCsv(header,csvPath)

	filepath.Walk(intrusionPath, func(ruleDirPath string, info os.FileInfo, err error) error {
		//ruleDirPath-->>rules/a/awstats_cmi_b_IPv6.xml
		if !util.IsDir(ruleDirPath) {
			//读取文件夹内每一个文件
			csvData := getXmlData(ruleDirPath)

			//for _,s:=range csvData{
			//	for _,v :=range s{
			//		fmt.Printf("===========>%s\n",v)
			//	}
			//}
			insertCsv(csvData,csvPath)
		}
		return nil
	})


}

/*
插入csc
	文件名 ExecDesc EngDesc Protocol  pro  payload
	a.xml value    value   value	 UDP  xxxx
*/
func insertCsv(csvC [][]string,csvPath string)bool  {
	csvFile,err:= util.FileMode(csvPath, 2)
	if err != nil {
		return false
	}

	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(csvFile) //创建一个新的写入文件流
	w.WriteAll(csvC)
	w.Flush()
	return true
}


func getXmlData(xmlNamePath string)  [][]string{
	file, err := os.Open(xmlNamePath)
	if err!=nil{
		fmt.Printf("open %s.xml err:%+v",xmlNamePath,err)
	}
	defer file.Close()
	bytes,_ := ioutil.ReadAll(file)
	stringReader := strings.NewReader(string(bytes))
	p := xml.NewDecoder(stringReader)


	headerElements := map[string]string{
		XmlName:path.Base(xmlNamePath),
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			element := token.Name.Local
			//是否包含3个键
			_, ExecDescOk := headerElements[ExecDesc]
			_, EngDescOk := headerElements[EngDesc]
			_, ProtocolOk := headerElements[Protocol]

			if ExecDescOk && EngDescOk && ProtocolOk{
				break
			}else {
				if element == ExecDesc{
					for _, attr := range token.Attr {
						valueKey := attr.Name.Local
						valueContent := attr.Value
						if valueKey == "value"{
							headerElements[ExecDesc] = valueContent
						}
					}
				}else if element == EngDesc{
					for _, attr := range token.Attr {
						valueKey := attr.Name.Local
						valueContent := attr.Value
						if valueKey == "value"{
							headerElements[EngDesc] = valueContent
						}
					}
				}else if element == Protocol{
					for _, attr := range token.Attr {
						valueKey := attr.Name.Local
						valueContent := attr.Value
						if valueKey == "value"{
							headerElements[Protocol] = valueContent
						}
					}
				}
			}
		}
	}

	//for k,v:=range headerElements{
	//	fmt.Printf("%s==>%s\n",k,v)
	//}

	finalData := [][]string{}
	stringReaderTmp := strings.NewReader(string(bytes))
	p = xml.NewDecoder(stringReaderTmp)

	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			element := token.Name.Local
			//fmt.Printf("element name: %s\n", element)
			//是否包含3个键
			for _, attr := range token.Attr {
				valueKey := attr.Name.Local//payPayload
				valueContent := attr.Value //payContent
				if valueKey == Payload{
					if valueContent != ""{
						protoMap := map[string]string{}

						//十进制转为
						//hex_data, _ := hex.DecodeString(valueContent)
						hex_data := util.Heco(valueContent)
						protoMap[element] = string(hex_data)
						csvContent := CreatCsvContent(headerElements,protoMap)

						finalData = append(finalData,csvContent)
					}
				}
			}
		}
	}
	return finalData
}

//构造csvcontent
func CreatCsvContent(headerEle map[string]string,protoEle map[string]string) []string {
	var k string
	var v string
	for k_,v_:=range protoEle{
		k = k_
		v =v_
	}

	tmpSlice := []string{
		headerEle[XmlName],
		headerEle[ExecDesc],
		headerEle[EngDesc],
		headerEle[Protocol],
		k,
		v,
	}
	return tmpSlice
}


