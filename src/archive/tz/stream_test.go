/*
@Time : 2020-12-17 18:13
@Author : wyf
@File : stream_test
@Software: GoLand
*/

package tz

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"testing"
)

func TestStream(t *testing.T) {
	Convey("Create a stream archive", t, func() {
		fpath := path.Join(os.TempDir(), "testdata/TestStream.tar.gz")
		err := os.MkdirAll(path.Dir(fpath), os.ModePerm)
		So(err, ShouldBeNil)
		defer os.Remove(fpath)

		fw, err := os.Create(fpath)
		So(err, ShouldBeNil)
		s := NewStreamArachive(fw)

		Convey("Stream a file", func() {
			f, err := os.Open("testdata/gophercolor16x16.png")
			So(err, ShouldBeNil)

			fi, err := f.Stat()
			So(err, ShouldBeNil)

			data := make([]byte, fi.Size())
			_, err = f.Read(data)
			So(err, ShouldBeNil)

			So(s.StreamFile("", fi, data), ShouldBeNil)
		})

		Convey("Stream a file with type directory", func() {
			f, err := os.Open("testdata")
			So(err, ShouldBeNil)

			fi, err := f.Stat()
			So(err, ShouldBeNil)

			So(s.StreamFile("", fi, nil), ShouldBeNil)
		})

		Convey("Stream a reader", func() {
			f, err := os.Open("testdata/gophercolor16x16.png")
			So(err, ShouldBeNil)

			fi, err := f.Stat()
			So(err, ShouldBeNil)

			So(s.StreamReader("", fi, f), ShouldBeNil)
		})
	})
}
