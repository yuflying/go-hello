// Copyright (c) 2018 Flyu, Inc.
//
// 服务启动参数配置 Created by flyu on 2018/11/12.
//

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h          bool
	ConfigFile string // ConfigFile 配置文件路径
	v          bool   // Version 是否输出版本信息
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&ConfigFile, "f", "configure.yaml", "config file path")
	flag.BoolVar(&v, "v", false, "show version and exit")

	// 改变默认的 Usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
welcome use hello
Usage: hello [-h] [-f configure file] [-v version]
	
Options:
`)
		flag.PrintDefaults()
	}
}

var (
	buildDate = "unknown"
	gitSHA    = "unknown"
)

// Version 返回版本信息
func Version() (date string, git string) {
	date = buildDate
	git = gitSHA
	return
}
