package str

import (
	"bytes"
	"strconv"
)

func Concat(str []string) string {
	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func ToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
