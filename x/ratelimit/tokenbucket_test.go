package ratelimit

import (
	"reflect"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	rl := New(1, time.Second)
	for i := 0; i < 2; i++ {

		if got := rl.Limit(); i == 0 && got {
			expected := false
			t.Errorf("wrong. expected: %v, got: %v", expected, got)
		}
		if got := rl.Limit(); i == 1 && !got {
			expected := false
			t.Errorf("wrong. expected: %v, got: %v", expected, got)
		}
	}

	//log.Printf("limit result: %v\n", rl.Limit())
}

func TestNew(t *testing.T) {
	type args struct {
		rate int
		per  time.Duration
	}
	tests := []struct {
		name string
		args args
		want *RateLimiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := New(tt.args.rate, tt.args.per); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRateLimiter_Limit(t *testing.T) {
	tests := []struct {
		name string
		rl   *RateLimiter
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.rl.Limit(); got != tt.want {
			t.Errorf("%q. RateLimiter.Limit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRateLimiter_UpdateRate(t *testing.T) {
	type args struct {
		rate int
	}
	tests := []struct {
		name string
		rl   *RateLimiter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.rl.UpdateRate(tt.args.rate)
	}
}

func TestRateLimiter_Undo(t *testing.T) {
	tests := []struct {
		name string
		rl   *RateLimiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.rl.Undo()
	}
}

func Test_unixNano(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := unixNano(); got != tt.want {
			t.Errorf("%q. unixNano() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
