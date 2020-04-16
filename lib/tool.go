package lib

import (
	"time"
)

func GetTimeStamp(data string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", data, loc)
	return tt.Unix()
}

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}