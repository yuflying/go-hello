// Copyright (c) 2018 Flyu, Inc.
//
// 生成uuid Created by flyu on 2018/10/25.
//

// twitter 雪花算法 (snowflake)
// 把时间戳,工作机器ID, 序列号组合成一个 64位 int
// 第一位置零, [2,42]这41位存放时间戳,[43,52]这10位存放机器id,[53,64]最后12位存放序列号
// https://www.cnblogs.com/Hollson/p/9116218.html

package uuid

import (
	"reflect"
	"time"
)

var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	serial        int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级)
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000
	SetMachineID(1000)
}

func SetMachineID(mid int64) {
	machineID = mid << 12 // 把机器ID左移12位，让出12位空间给序列号使用
}

func GetSnowFlakeID() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000
	if curTimeStamp > lastTimeStamp {
		serial = 0
		lastTimeStamp = curTimeStamp
		ts := curTimeStamp & 0x1FFFFFFFFFF
		ts <<= 22
		return ts | machineID | serial
	} else if curTimeStamp == lastTimeStamp {
		serial++
		if serial > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			lastTimeStamp = curTimeStamp
			serial = 0
		}
		ts := curTimeStamp & 0x1FFFFFFFFFF
		ts <<= 22
		return ts | machineID | serial
	} else {
		return 0
	}
}

//去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func New() int64 {
	return GetSnowFlakeID()
}
