// Copyright (c) 2018 Flyu, Inc.
//
// 默认配置 Created by flyu on 2018/08/19.
//

package constant

var (
	// Debug 开启debug模式
	Debug = true
	// Port 默认启动端口号
	Port = 10011
)

// DataBase 默认数据库配置
var DB = struct {
	IP          string
	Port        int
	User        string
	Password    string
	Name        string
	MaxIdleTime int // 连接池最大空闲时间
	MaxIdle     int // 最大空闲数量
	Pool        int // 最大连接池数量
	MaxOverflow int // 最大上限
}{
	IP:          "127.0.0.1",
	Port:        3306,
	User:        "root",
	Password:    "flyu",
	Name:        "flyu",
	MaxIdleTime: 1800,
	MaxIdle:     10,
	Pool:        10,
	MaxOverflow: 10,
}

// Redis 默认配置
var Redis = struct {
	IP          string
	Port        int
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int64 //ms
	Timeout     struct {
		Connect int64
		Write   int64
		Read    int64
	}
}{
	IP:          "127.0.0.1",
	Port:        6379,
	DB:          0,
	Password:    "",
	MaxIdle:     10,
	MaxActive:   100,
	IdleTimeout: 300,
	Timeout: struct {
		Connect int64
		Write   int64
		Read    int64
	}{Connect: 10, Write: 10, Read: 10},
}
