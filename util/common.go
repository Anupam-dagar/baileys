package util

import (
	"github.com/Anupam-dagar/baileys/constant"
	"time"
)

func GetNowTimeMillis() string {
	return time.Now().Format(constant.TimeFormatMillis)
}
