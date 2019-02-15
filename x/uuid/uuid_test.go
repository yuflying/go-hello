package uuid

import (
	"testing"
)

// go test -v *.go -test.run TestSetMachineID
func TestSetMachineID(t *testing.T) {
	type args struct {
		mid int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{11}, 45056},
		{"2", args{22}, 90112},
	}
	for _, tt := range tests {
		SetMachineID(tt.args.mid)
		if machineID != tt.want {
			t.Errorf("%q. SetMachineID(tt.args.mid) = %v, want %v", tt.name, machineID, tt.want)
		}
	}
}

// go test -v *.go -test.run TestGetSnowFlakeID
func TestGetSnowFlakeID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"1", 6463756154912657409},
		{"2", 6463756154912657410},
	}
	for _, tt := range tests {
		if got := GetSnowFlakeID(); got == tt.want {
			t.Errorf("%q. GetSnowFlakeID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// go test -v *.go -test.run TestDuplicate
func TestDuplicate(t *testing.T) {
	// 设置一个机器标识，如IP编码,防止分布式机器生成重复码
	SetMachineID(192168100101)
	var ids = make([]int64, 0)
	for i := 0; i < 10000000; i++ {
		id := GetSnowFlakeID()
		ids = append(ids, id)
	}

	tests := []struct {
		name    string
		args    interface{}
		wantRet int
	}{
		{"1", ids, 10000000},
	}
	for _, tt := range tests {
		if gotRet := len(Duplicate(tt.args)); gotRet != tt.wantRet {
			t.Errorf("%q. Duplicate() = %v, want %v", tt.name, gotRet, tt.wantRet)
		}
	}
}

// go test -v *.go -test.run TestNew
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"1", 6463756154912657409},
		{"2", 6463756154912657410},
	}
	for _, tt := range tests {
		if got := New(); got == tt.want {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
