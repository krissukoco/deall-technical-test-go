package tests

import (
	"time"

	"github.com/google/uuid"
)

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func Now() int64 {
	return time.Now().UnixMilli()
}

func RandomId() string {
	return uuid.New().String()
}
