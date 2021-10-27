/*
@Time : 2021/10/25 14:35
@Author : wyf
@File : bytes_test
@Software: GoLand
*/

package bytes

import (
	"fmt"
	"testing"
)



func TestBytesBufferWriteString(t *testing.T) {
	strs := []string{
		"safly",
		"xinmu",
	}
	ret := BytesBufferWriteString(strs...)
	fmt.Println(ret)
}
