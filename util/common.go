package util

import (
	"errors"
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

func ReadTag(modelStruct any, fieldName string, tag string) (string, error) {
	field, ok := reflect.TypeOf(modelStruct).Elem().FieldByName(fieldName)
	if !ok {
		return "", errors.New("field not found")
	}
	return field.Tag.Get(tag), nil
}
