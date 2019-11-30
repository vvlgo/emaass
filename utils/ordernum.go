package utils

import (
	"math/rand"
	"strconv"
	"time"
)

/*
   GetOrderNum  计算订单号，订单号格式yyyymmddDDmmss+5位随机数
*/
func GetOrderNum() string {
	ntime := time.Now()
	s := ntime.Format("20060102150405")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result = ""
	for i := 0; i < 5; i++ {
		result += strconv.Itoa(r.Intn(10))
	}
	s = s + result
	return s
}
