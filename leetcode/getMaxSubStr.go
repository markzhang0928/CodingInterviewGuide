package main

import "bytes"

func getMaxSubStr2(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var buf bytes.Buffer
	maxLen := 0    // 公共子串的最大长度
	maxLenEnd := 0 //

	for i := 0; i < len1+len2; i++ {
		s1begin, s2begin := 0, 0
		tmpMaxLen := 0
		if i < len1 {
			s1begin = len1 - i
		} else {
			s2begin = i - len1
		}
		var j = 0
		for j = 0; (s1begin+j < len1) && (s2begin+j < len2); j++ {
			if str1[s1begin+j] == str2[s2begin+j] {
				tmpMaxLen++
			} else {
				if tmpMaxLen > maxLen {
					maxLen = tmpMaxLen
					maxLenEnd = s1begin + j
				} else {
					tmpMaxLen = 0
				}
			}
		}
		if tmpMaxLen > maxLen {
			maxLen = tmpMaxLen
			maxLenEnd = s1begin + j
		}
	}
	for i := maxLenEnd - maxLen; i < maxLenEnd; i++ {
		buf.WriteByte(str1[i])
	}
	return buf.String()
}
