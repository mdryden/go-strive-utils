package stringutils

func MaskLeft(s string, keepCount int, maskChar rune) string {
	rs := []rune(s)
	for i := 0; i < len(rs)-keepCount; i++ {
		rs[i] = maskChar
	}
	return string(rs)
}
