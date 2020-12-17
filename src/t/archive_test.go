/*
@Time : 2020-12-16 23:56
@Author : wyf
@File : archive_test
@Software: GoLand
*/

package t

import (
	"fmt"
	"goapi/src/libs"
	"testing"
)

func TestTarFile(t *testing.T) {
	err := libs.TarFile("file")

	fmt.Printf("%+v",err)
}




func TestUnTarFile(t *testing.T) {
	err := libs.UnTarFile("file.tar")

	fmt.Printf("%+v",err)
}
func TestTarFiles(t *testing.T) {
	err := libs.TarFiles("files")

	fmt.Printf("%+v",err)
}

func TestUnTarFiles(t *testing.T) {
	err := libs.UnTarFiles("files.tar")

	fmt.Printf("%+v",err)
}

