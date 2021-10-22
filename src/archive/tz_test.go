/*
@Time : 2020-12-18 00:16
@Author : wyf
@File : tz_test
@Software: GoLand
*/

package archive

import (
	"testing"
)

//打包多个文件
func TestTarFiles(t *testing.T) {
	srcs := []string{"file1", "file2"}
	Tar(srcs, "files.tar.gz", true)
}

//打包一个文件
func TestTarFile(t *testing.T) {
	srcs := []string{"file1"}
	Tar(srcs, "file1.tar.gz", true)
}

//解压一个文件，至当前
func TestUnTarFile(t *testing.T) {
	UnTar("file1.tar.gz", "")
}

//解压一个文件至目录
func TestUnTarFileToDir(t *testing.T) {
	UnTar("file1.tar.gz", "untar")
}

//解压多个文件，至当前
func TestUnTarFiles(t *testing.T) {
	UnTar("files.tar.gz", "")
}

//解压多个文件，至目录
func TestUnTarFilesToDir(t *testing.T) {
	UnTar("files.tar.gz", "untars")
}
