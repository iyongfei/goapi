/*
@Time : 2020-12-17 00:29
@Author : wyf
@File : t
@Software: GoLand
*/

package main

import (
	"fmt"
	"os"
)

func main()  {
	fi, statErr := os.Stat("file")
	if statErr!=nil{
		fmt.Printf("statErr:%+v",statErr)
	}

	fmt.Println(fi.Name())
}
