// Copyright (c) 2018 Flyu, Inc.
//
// redis Created by flyu on 2018/08/25.
//

package comm

import (
	"testing"
	_ "x/mysql"
)

// go test -v *.go -test.run TestInitDB
func TestInitDB(t *testing.T) {
	type args struct {
		conf *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{
			&Config{
				DB: Database{
					IP:          "127.0.0.1",
					Port:        3306,
					User:        "root",
					Password:    "root",
					Name:        "sso",
					MaxIdleTime: 1800,
					MaxIdle:     10,
					Pool:        10,
					MaxOverflow: 10,
				},
			},
		}, false},
	}
	for _, tt := range tests {
		if err := InitDB(tt.args.conf); (err != nil) != tt.wantErr {
			t.Errorf("%q. InitDB() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
