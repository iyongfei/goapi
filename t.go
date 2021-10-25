/*
@Time : 2020-12-17 00:29
@Author : wyf
@File : t
@Software: GoLand
*/

package main

import "io"

//type Inter interface {
//	say()
//
//}

type S struct {
	rd           io.Reader
	err          error
}

func main() {


}

func (b *S) Read(p []byte) (n int, err error) {
	return 0,nil
}

func Func(in io.Reader)  {




}
