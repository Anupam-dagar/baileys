package util

import (
	"baileys/constant"
	"time"
)

func GetNowTimeMillis() string {
	return time.Now().Format(constant.TimeFormatMillis)
}
