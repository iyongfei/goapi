/*
@Time : 2020-12-18 00:16
@Author : wyf
@File : tz_test
@Software: GoLand
*/

package archive

import (
	"goapi/src/archive/tz"
	"testing"
)

//打包多个文件
func TestTarFiles(t *testing.T) {
	srcs := []string{"file1", "file2"}
	tz.Tar(srcs, "files.tar.gz", true)
}

//打包一个文件
func TestTarFile(t *testing.T) {
	srcs := []string{"file1"}
	tz.Tar(srcs, "file1.tar.gz", true)
}

//解压一个文件，至当前
func TestUnTarFile(t *testing.T) {
	tz.UnTar("file1.tar.gz", "")
}

//解压一个文件至目录
func TestUnTarFileToDir(t *testing.T) {
	tz.UnTar("file1.tar.gz", "untar")
}

//解压多个文件，至当前
func TestUnTarFiles(t *testing.T) {
	tz.UnTar("files.tar.gz", "")
}

//解压多个文件，至目录
func TestUnTarFilesToDir(t *testing.T) {
	tz.UnTar("files.tar.gz", "untars")
}
