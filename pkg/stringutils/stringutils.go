package stringutils

import (
	"fmt"
	"strings"
)

func JoinArray[T interface{}](values []T, delimiter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), ","), "[]")
}
