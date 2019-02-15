// Copyright (c) 2018 Flyu, Inc.
//
// 全局配置 Created by flyu on 2018/08/19.
//

package comm

import (
	"embrace/constant"
	"io/ioutil"
	"unsafe"
	"x/yaml"
)

var Conf *Config

type Config struct {
	Port   int      `yaml:"port"`
	FlyuDB Database `yaml:"flyu_db"`
}

// Database 数据库配置
type Logs struct {
	Time   bool `yaml:"time"`
	Debug  bool `yaml:"debug"`
	Trace  bool `yaml:"trace"`
	Colors bool `yaml:"colors"`
	Pid    bool `yaml:"pid"`
}

// Database 数据库配置
type Database struct {
	IP          string `yaml:"ip"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"pass"`
	Name        string `yaml:"name"`
	MaxIdleTime int    `yaml:"max_idle_time"`
	MaxIdle     int    `yaml:"max_idle"`
	Pool        int    `yaml:"pool"`
	MaxOverflow int    `yaml:"max_overflow"`
}

// NewConf 根据配置文件生成配置信息
func NewConfig(data []byte) (conf *Config, err error) {
	conf = &Config{}
	err = yaml.Unmarshal(data, conf)
	return
}

// InitConf 读取配置文件或全局变量，生成全局配置信息
func LoadConfig(path string) {
	if Conf == nil {
		data, err := ioutil.ReadFile(path)
		if err != nil { // 从全局变量里取
			Conf = &Config{}
			Conf.Port = constant.Port
			Conf.FlyuDB = *(*Database)(unsafe.Pointer(&constant.DB))
			Log.Error("[load config from file failed, Error: %s]", err.Error())
			Log.Trace("[default conf: %+v]", Conf)
			return
		}
		if Conf, err = NewConfig(data); err != nil {
			Log.Fatal("[解析配置文件出错,Error: %s]", err.Error())
		}
	}
	Log.Debug("[load config from file succeed, conf: %+v]", Conf)
}
