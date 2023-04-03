package util

import (
	"github.com/Anupam-dagar/baileys/constant"
	"reflect"
	"strings"
	"time"
)

func GetNowTimeMillis() string {
	return time.Now().Format(constant.TimeFormatMillis)
}

func SplitStringFromBack(str, delimiter string) (string, string) {
	lastIndex := strings.LastIndex(str, delimiter)
	if lastIndex == -1 {
		return "", str
	}
	return str[:lastIndex], str[lastIndex+len(delimiter):]
}

func ReadTag(modelStruct any, fieldName string, tag string) string {
	field, ok := reflect.TypeOf(modelStruct).Elem().FieldByName(fieldName)
	if !ok {
		panic("Field not found")
	}
	return field.Tag.Get(tag)
}
