package main

import (
	"embrace/comm"
	"flag"
	"fmt"
	"os"
)

// 初始化全局配置
func initConfig() {

	// 初始化日志库
	comm.InitLogger()

	// 初始化命令行参数
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if v {
		fmt.Println("version: 1.0.0_01")
		os.Exit(0)
	}

	// 加载全局配置文件
	comm.LoadConfig(ConfigFile)

	// 初始化数据库
	comm.InitDB()
}
