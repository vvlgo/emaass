package utils

import "strconv"

/*
提取员工编号的数字,并生成新的编号
*/
func GetEmpnum(str string, i int) string {
	runes := []rune(str)
	var s string
	var s2 string
	var num int
	for k, v := range runes {
		if v > 48 && v <= 57 {
			s = string(runes[k:])
			num, _ = strconv.Atoi(s)
			n := len(strconv.Itoa(num+i)) - len(s)
			if num+i < 10 {
				s2 = string(runes[:k])
			}
			if num+i >= 10 {
				s2 = string(runes[:k-n])
			}
			break
		}
	}
	newNum := s2 + strconv.Itoa(num+i)
	return newNum
}
