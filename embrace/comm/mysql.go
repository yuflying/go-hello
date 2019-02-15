// Copyright (c) 2018 Flyu, Inc.
//
// redis Created by flyu on 2018/08/25.
//

package comm

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	_ "x/mysql" // init mysql
)

var (
	FlyuDB    *sqlx.DB // flyu数据库
	ShandjjDB *sqlx.DB // 闪电降价数据库
)

func InitDB() {
	var err error

	// 初始化闪电降价数据库
	if FlyuDB == nil {
		if FlyuDB, err = newDB(Conf.FlyuDB); err != nil {
			Log.Fatal("init flyu db failed, Error: %s", err)
		}
	}
}

// 新建DB
func newDB(conf Database) (db *sqlx.DB, err error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8",
		conf.User,
		conf.Password,
		conf.IP,
		conf.Port,
		conf.Name,
	)

	// 设置默认值
	if conf.MaxIdle < 0 {
		conf.MaxIdle = 10
	}
	if conf.MaxIdleTime < 0 {
		conf.MaxIdleTime = 1800
	}
	if conf.MaxOverflow < 0 {
		conf.MaxOverflow = 20
	}

	db, err = sqlx.Open("mysql", url)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(conf.MaxIdle)
	db.SetConnMaxLifetime(time.Duration(conf.MaxIdleTime) * time.Second)
	db.SetMaxOpenConns(conf.MaxOverflow + 10)
	return db, db.Ping()
}
