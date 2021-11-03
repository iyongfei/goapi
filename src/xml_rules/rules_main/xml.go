/**
find . -name "*IPv6.xml"  -maxdepth 2 -delete
ls -l |grep "^-"|wc -l
ls -l |grep "^d"|wc -l
 */

/**

#icmp:          4020000~4020999
#floodICMPDFwhenfragsneeded.xml
alert tcp any any -> any any (msg:"123"; content:"456"; sid:4035000; rev:1; metadata:created_at 2021_10_26, updated_at 2021_10_26;)
 */
package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"goapi/src/xml_rules/util"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)
//go run rules_xml_rules/xml.go
const (
	baseXmlDir = "excel_rules"
)

const (
	XmlName = "XmlName"
	ExecDesc = "ExecDesc"
	Payload = "Payload"
)
var (
	intrusionPath = "rules/bypass"
	ruleDirName=""
	rulePath = ""
	//rulePath = "v-bypass.rules"
	startIndex = 4020000
)
func init()  {
	fmt.Println("in..................")
	var ruleDir string
	flag.StringVar(&ruleDir, "v", "", "vid")
	flag.Parse()

	intrusionPath = ruleDir


	if !util.Exists(baseXmlDir){
		_= os.Mkdir(baseXmlDir, os.ModePerm)
	}

	/**
	初始化文件路径
	 */
	segIndex := strings.Index(intrusionPath,"/")
	//bypass(xml目录)
	ruleDirName= intrusionPath[segIndex+1:]
	//v-bypass.rules
	rulePath = fmt.Sprintf("v-%s.rules",ruleDirName)
	fmt.Println("ruleDirName----->",ruleDirName,",rulePath----->",rulePath)
	/**
	初始化serialnum
	 */
	serialStr := util.RuleHeaderMap[ruleDirName]
	right := strings.Split(serialStr,":")[1]

	startSearial := strings.Split(right,"~")[0]
	startIndex,_ = strconv.Atoi(startSearial)
	fmt.Println("startIndex----->",startIndex)
}

func main()  {

	//获取rules/bypass下的所有xml
	filepath.Walk(intrusionPath, func(ruleDirPath string, info os.FileInfo, err error) error {
		//ruleDirPath-->>rules/bypass/awstats_cmi_b_IPv6.xml
		if !util.IsDir(ruleDirPath) {
			//读取文件夹内每一个xml
			 getXmlData(ruleDirPath)
		}
		return nil
	})

	//插入头
	insertData(rulePath,util.RuleHeaderMap[ruleDirName]+"\n\n")

	for _,ruleRet:=range ruleRets{
		//构造字符串插入
		xName := ruleRet.XmlName
		proto := ruleRet.Protocol
		execDesc := ruleRet.ExecDesc
		serialNum := ruleRet.SerialNum


		str := fmt.Sprintf(`#%s\nalert %s any any -> any any (msg:"%s"; content:\"456\"; sid:%d; rev:1; metadata:created_at 2021_10_26, updated_at 2021_10_26;)\n\n`,xName,proto,execDesc,serialNum)
		fmt.Println(str)
		insertData(rulePath,str)
	}
}
/**
#icmp:          4020000~4020999
#floodICMPDFwhenfragsneeded.xml
alert tcp any any -> any any (msg:"123"; content:"456"; sid:4035000; rev:1; metadata:created_at 2021_10_26, updated_at 2021_10_26;)
*/

func insertData(fpath string,str string) bool {
	fPath := baseXmlDir + "/" + fpath
	csvFile,err:= util.FileMode(fPath, 2)
	if err != nil {
		return false
	}
	writer := bufio.NewWriter(csvFile)
	writer.WriteString(str)
	writer.Flush()

	return true
}

func getXmlData(xmlNamePath string)  {
	file, err := os.Open(xmlNamePath)
	if err!=nil{fmt.Printf("open %s.xml err:%+v",xmlNamePath,err)}
	defer file.Close()
	bytes,_ := ioutil.ReadAll(file)
	stringReader := strings.NewReader(string(bytes))
	p := xml.NewDecoder(stringReader)

	headerElements := map[string]string{
		XmlName:path.Base(xmlNamePath),
	}

	headerProto := map[string]string{
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			element := token.Name.Local
			if element == ExecDesc{
				for _, attr := range token.Attr {
					valueKey := attr.Name.Local
					valueContent := attr.Value
					if valueKey == "value"{
						//headerElements[ExecDesc] = valueContent

						s1 := strings.Replace(valueContent,`\`,`\\`,-1)
						s2 := strings.Replace(s1,`"`,`\"`,-1)
						s3 := strings.Replace(s2,`;`,`\;`,-1)
						s4 := strings.Replace(s3,`:`,`\:`,-1)
						headerElements[ExecDesc] = s4
					}
				}
			}
			/**
			<IP Version="04" InternetHeaderLength="05" TypeofService="00" TotalLength="05DC" Identification="3333" FragmentOffset="0000" TTL="40" Protocol="11" HeaderChecksum="B816" srcIP="$sourceIP" destIP="$destIP" options="" CE="false" DF="false" MF="true">
			        <UDP srcPort="$sourcePort" destPort="$destPort" checksum="C6FE" length="0BF8" Payload="B0EEF0110000000000000002000186A0000000020000000500000000000000000000000000000000000186A0000000020000000000000BB8414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141"/>
			</IP>
			 */
			for _, attr := range token.Attr {
				valueKey := attr.Name.Local//payPayload
				//valueContent := attr.Value //payContent
				if valueKey == Payload{
					//if valueContent != ""{
					headerProto[element] = element
					//}
				}
			}
		}
	}
	for k,_:=range headerProto{
		ruleRet := &RuleRet{
			XmlName: headerElements[XmlName],
			ExecDesc: headerElements[ExecDesc],
			SerialNum: startIndex,
			Protocol: k,
		}
		ruleRets = append(ruleRets,ruleRet)
		startIndex += 1
	}
}



type RuleRet struct {
	XmlName string
	Protocol string
	ExecDesc string
	SerialNum int
}
var ruleRets = []*RuleRet{}
