package tests

import "time"

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func Now() int64 {
	return time.Now().UnixMilli()
}
