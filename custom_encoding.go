package urlshortener

import (
	"strconv"
)

func EncodeBase36(num int64) (string, error) {
	return strconv.FormatInt(num, 36), nil
}

func DecodeBase36(s string) (int64, error) {
	return strconv.ParseInt(s, 36, 64)
}
