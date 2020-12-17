/*
@Time : 2020-12-16 23:55
@Author : wyf
@File : archive
@Software: GoLand
*/

package libs

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func TarFile(srcFile string)  error{
	//创建desFileErr
	var desFileName = fmt.Sprintf("%s.tar",srcFile)

	desFile, desFileErr:= os.Create(desFileName)

	if desFileErr!=nil{
		fmt.Printf("desFileErr:%+v",desFileErr)
		return desFileErr
	}
	defer func() {
		if err := desFile.Close(); err != nil {
			fmt.Printf("%+v",err)
		}
	}()

	//newWriter for destFile
	tw := tar.NewWriter(desFile)


	defer func() {
		if err := tw.Close(); err != nil {
			fmt.Printf("%+v",err)
		}
	}()
	//源文件是否合法
	fi, statErr := os.Stat(srcFile)
	if statErr!=nil{
		fmt.Printf("statErr:%+v",statErr)
		return statErr
	}

	hdr, headerErr := tar.FileInfoHeader(fi, "")
	if headerErr!=nil{
		fmt.Printf("headerErr:%+v",headerErr)
		return headerErr
	}


	writerErr := tw.WriteHeader(hdr)
	if writerErr!=nil{
		fmt.Printf("writerErr:%+v",writerErr)
		return writerErr
	}
	//打开源文件
	fr, openErr := os.Open(srcFile)
	if openErr!=nil{
		fmt.Printf("openErr:%+v",openErr)
		return openErr
	}

	defer func() {
		if err := fr.Close(); err != nil {
			fmt.Printf("%+v",err)
		}
	}()

	//复制内容到destFile
	written, err := io.Copy(tw, fr)
	log.Printf("共写入了 %d 个字符的数据\n",written)
	return err
}

func UnTarFile(srcFile string)  error{
	//打开压缩文件
	fr, openSrcFileErr := os.Open(srcFile)
	defer fr.Close()

	if openSrcFileErr!=nil{
		fmt.Printf("openSrcFileErr:%+v",openSrcFileErr)
		return openSrcFileErr
	}
	//newReader for srcFile
	tr := tar.NewReader(fr)

	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next(){
		// 获取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, createErr := os.Create(fi.Name())

		if createErr!=nil{
			fmt.Printf("createErr:%+v",createErr)
			return createErr
		}

		// 将 tr 写入到 fw
		n, copyErr := io.Copy(fw, tr)
		log.Printf("解包： %s 到 %s ，共处理了 %d 个字符的数据。", srcFile,fi.Name(),n)

		if copyErr!=nil{
			fmt.Printf("copyErr:%+v",copyErr)
			return copyErr
		}
		os.Chmod(fi.Name(),fi.Mode().Perm())
		fw.Close()
	}

	return  nil
}



func TarFiles(srcFile string)  error{
	var desFileName = fmt.Sprintf("%s.tar",srcFile)
	fw, err := os.Create(desFileName)
	if err != nil {
		return err
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)

	defer tw.Close()


	return filepath.Walk(srcFile, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		log.Printf("成功打包 %s ，共写入了 %d 字节的数据\n", fileName, n)

		return nil
	})

}

func UnTarFiles(srcFile string)  error{
	// 打开准备解压的 tar 包
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()

	// 将打开的文件先解压
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()


	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}

		switch hdr.Typeflag {
		case tar.TypeDir:
			if b := ExistDir(hdr.Name); !b {
				if err := os.MkdirAll(hdr.Name, 0775); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(hdr.Name, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			n, err := io.Copy(file, tr)
			if err != nil {
				return err
			}
			fmt.Printf("成功解压： %s , 共处理了 %d 个字符\n", hdr.Name, n)

			file.Close()
		}
	}

	return nil

}
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

