package crypt

import (
	"testing"
	"time"
)

var TokenTimeout = time.Second * 30
var PrivKey = "12345678900987654321123456789098"

// go test -count=1 -v *.go -test.run TestNewToken
func TestNewToken(t *testing.T) {
	type args struct {
		key string
		m   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{PrivKey, map[string]interface{}{
			"sso_uuid":   "6461200898559332352",
			"encryption": 1,
			"exp":        time.Now().Add(10000 * time.Hour).Unix(),
		}}, "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbmNyeXB0aW9uIjoxLCJleHAiOjE1NzY4OTQ2ODEsInNzb191dWlkIjoiNjQ2MTIwMDg5ODU1OTMzMjM1MiJ9.acwDrnn9qHTpssqYSy7J14yfXcavPq6QjKVbWyzxT5_zGzNeT2UL43LJ4GhEgr7im1CjQy6hTFeG0PL8iAbN5Q"},
	}

	for _, tt := range tests {
		if got := NewToken(tt.args.key, tt.args.m); got != tt.want {
			t.Errorf("%q. NewToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// go test -v -count=1 *.go -test.run TestParseToken
func TestParseToken(t *testing.T) {
	type args struct {
		tokenStr string
		key      string
	}
	p := args{
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbmNyeXB0aW9uIjoxLCJleHAiOjE1NzY4OTQ2ODEsInNzb191dWlkIjoiNjQ2MTIwMDg5ODU1OTMzMjM1MiJ9.acwDrnn9qHTpssqYSy7J14yfXcavPq6QjKVbWyzxT5_zGzNeT2UL43LJ4GhEgr7im1CjQy6hTFeG0PL8iAbN5Q", PrivKey,
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", p, false},
	}
	for _, tt := range tests {
		got, err := ParseToken(tt.args.tokenStr, tt.args.key)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ParseToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		t.Logf("got: %+v", got)
	}
}

// go test -count=1 *.go -test.run TestValidateToken
func TestValidateToken(t *testing.T) {
	type args struct {
		tokenStr string
		key      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{
			"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbmNyeXB0aW9uIjoxLCJleHAiOjE1NzY4OTA1MDcsInNzb191dWlkIjoiOTk5OTk5OTk5OTk5OTk5OTk5OSJ9.USjn4UoFb3GuNzYkV_kZUv4IXVsO8WPNE9bl1cteKEujAoaG-8m55gq5QjZGtaghbDetWkGShfeZuzcrOTmNWQ", PrivKey,
		}, false},
		{"2", args{
			"fdsfsfdsflds", PrivKey,
		}, true},
	}
	for _, tt := range tests {
		if err := ValidateToken(tt.args.tokenStr, tt.args.key); (err != nil) != tt.wantErr {
			t.Errorf("%q. ValidateToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

// go test -count=1 *.go -test.run TestEncryptPwd
func TestEncryptPwd(t *testing.T) {
	type args struct {
		pwd  string
		salt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"123456", "hehe"}, "00f37da9922c1ff4ad808fc9775f1729"},
		{"2", args{"abcdef", "hehe"}, "a8d90e73f5e281989ea2ce127fee826b"},
		{"3", args{"ssssss", "hehe"}, "c5b4a42581c2cb174dc7b4fbda9f5994"},
	}
	for _, tt := range tests {
		if got := EncryptPwd(tt.args.pwd, tt.args.salt); got != tt.want {
			t.Errorf("%q. EncryptPwd() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
