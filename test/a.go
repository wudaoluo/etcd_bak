package main

import (
	"fmt"
	"time"
)

func main() {
	//时间戳转日期
	timeLayout := "20060102-15:04:05"
	dataTimeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
	//fmt.Println(strconv.FormatInt(time.Now().Unix(),10))
	//fmt.Println(path.Join())
}
