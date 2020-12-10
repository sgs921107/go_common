/*************************************************************************
	> File Name: funcs.go
	> Author: xiangcai
	> Mail: xiangcai@gmail.com
	> Created Time: 2020年12月09日 星期三 19时39分01秒
 ************************************************************************/

package gcommon

import (
	"bytes"
	"io"
	"net/url"
	"strings"
	"time"
)

type TimeStampFlag int

const (
	SECOND TimeStampFlag = iota
	MILLISECOND
	MICROSECOND
	NANOSECOND
)

// convert type map to []byte
func MapToBytes(data map[string]string) []byte {
	reader := MapToReader(data)
	return ReaderToBytes(reader)
}

// convert type map to io.reader
func MapToReader(data map[string]string) io.Reader {
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	return strings.NewReader(form.Encode())
}

// convert type io.reader to []byte
func ReaderToBytes(reader io.Reader) []byte {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)
	return buffer.Bytes()
}

// convert type io.reader to string
func ReaderToString(reqBody io.Reader) string {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reqBody)
	return buffer.String()
}

/*
获取时间戳
接收一个整形 0-3
0-秒, 1-毫秒, 2-微妙, 3-纳秒
*/
func TimeStamp(flag int) int64 {
	now := time.Now()
	switch TimeStampFlag(flag) {
	case SECOND:
		return now.Unix()
	case MILLISECOND:
		return now.UnixNano() / 1e6
	case MICROSECOND:
		return now.UnixNano() / 1e3
	case NANOSECOND:
		return now.UnixNano()
	default:
		return 0
	}
}