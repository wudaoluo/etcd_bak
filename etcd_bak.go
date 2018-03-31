package main

import (
	"flag"
	"os/signal"
	"syscall"
	"os"
	"./utils"
	"./core"
	"time"
)

type argStruct struct {
	version bool
	debug bool
	logdir string
	etcdaddr string
	interval int64
	confdir string

}

var arg = new(argStruct)

func init() {
	flag.BoolVar(&arg.version,"version",false,"print version")
	flag.BoolVar(&arg.debug,"debug",true,"open debug default false")
	flag.StringVar(&arg.logdir,"logdir","log","日志目录")
	flag.StringVar(&arg.etcdaddr,"etcdaddr","http://192.168.222.140:2379","etcd连接地址")
	flag.Int64Var(&arg.interval,"interval",3600,"每次检查指纹间隔时间")
	flag.StringVar(&arg.confdir,"confdir","conf","配置保存目录")
	flag.Parse()
}


func main() {
	if arg.version {
		utils.PrintVersion()
		os.Exit(0)
	}

	w := core.NewWorker("etcd_dir",[]string{"http://10.81.46.99:2379","http://10.81.178.198:2379","http://10.81.178.99:2379"})

	//删除本地的值,重新更新所有值

	//如果获取etcd 失败,则sleep 60s在获取
	for {
		if w.GetRoot() {
			break
		}
		time.Sleep(60*time.Second)
	}


	w.SyncDir()

	w.Watch("/")

	//初始化日志
	utils.LogInit(true,arg.logdir)

	waitSignal()
}


//阻塞，只有执行信号才执行
func waitSignal() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-osSignals
}