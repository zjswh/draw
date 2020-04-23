package lib

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

func GetTimeStamp(data string) int64 {
	arr := strings.Fields(data)
	length := len(strings.Split(arr[0],"-"))
	len2 := 0
	if arr[1] != "" {
		len2  += len(strings.Split(arr[1],":"))
	}
	length += len2
	fields := [6]string{"2006","01","02","15","04","05"}
	var layout string
	switch length {
		case 1: layout = fields[0]
		case 2: layout = fields[0] + "-" +  fields[1]
		case 3: layout = fields[0] + "-" +  fields[1] + "-" + fields[2]
		case 4: layout = fields[0] + "-" +  fields[1] + "-" + fields[2] + " " + fields[3]
		case 5: layout = fields[0] + "-" +  fields[1] + "-" + fields[2] + " " + fields[3] + ":" +  fields[4]
		case 6: layout = fields[0] + "-" +  fields[1] + "-" + fields[2] + " " + fields[3] + ":" +  fields[4] + ":" + fields[5]
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation(layout, data, loc)
	return tt.Unix()
}

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func Md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
