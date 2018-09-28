package base

import (
	"fmt"
	"errors"
	"strconv"
)

func StrVal(data interface{}) string {
	switch v := data.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 64)
	default:
		return fmt.Sprintf("%v", data)
	}
}

func IntVal(data interface{}) (int, error) {
	switch v := data.(type) {
	case string:
		return strconv.Atoi(v)
	default:
		return 0, errors.New("unknown type")
	}
}