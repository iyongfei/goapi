/*
@Time : 2020-12-18 00:16
@Author : wyf
@File : tz_test
@Software: GoLand
*/

package archive

import (
	"bytes"
	"fmt"
	"github.com/yeka/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestZipyeka(t *testing.T)  {
	contents := []byte("Hello World")
	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	w, err := zipw.Encrypt(`test.txt`, `02bb9d4d2d62a1cb9e89989536bccea2`, zip.AES128Encryption)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
}


func TestUnZipyeka(t *testing.T)  {
	r, err := zip.OpenReader("encrypted.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword("02bb9d4d2d62a1cb9e89989536bccea2")
		}

		r, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		fmt.Printf("Size of %v: %v byte(s),str:%s\n", f.Name, len(buf),string(buf))
	}
}
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
