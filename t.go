/*
@Time : 2020-12-17 00:29
@Author : wyf
@File : t
@Software: GoLand
*/

package main

import (
	"bufio"
	"fmt"
	"strings"
)


func main() {
	s := strings.NewReader(strings.Repeat("a", 40) + "s|ss")
	r := bufio.NewReader(s)
	token, err := r.ReadBytes('|')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q\n", token)

}

