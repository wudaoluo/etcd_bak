package core

import (
	"time"
	"strconv"
)



var process_start_time string

//虚拟目录,每次启动先将etcd 数据放到这里,在删除 etcd_dir,再mv 这个目录 etcd_dir
var dir_timestamps = strconv.FormatInt(time.Now().Unix(),10)

func init() {
	process_start_time = sys_time()
}


//返回当前系统时间
func sys_time() string {
	time_format := "20060102-15:04:05"
	return time.Unix(time.Now().Unix(), 0).Format(time_format)
}