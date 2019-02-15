// Copyright (c) 2018 Flyu, Inc.
//
// log Created by flyu on 2018/08/19.
//

package comm

import (
	"x/logger"
)

var Log *logger.Logger

func InitLogger() {
	Log = logger.NewStdLogger(true, true, true, true, true)
}
