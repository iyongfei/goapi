/**
如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段
如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段
如果某个struct字段的tag定义了中含有",attr"，那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段
如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段

 */

package main

import "fmt"

func main()  {

	fmt.Println("sssssss")

}

