package main

import (
	"testing"
)

func Test_genUniqueID(t *testing.T) {
	type args struct {
		mainStart int64
		increment int
		now       int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "start time is present",
			args: args{
				mainStart: 1593008100,
				increment: 0,
				now:       1593008100,
			},
			want: "000005ef35fe4001",
		}, {
			name: "start time is present + 1",
			args: args{
				mainStart: 1593008100,
				increment: 1,
				now:       1593008100,
			},
			want: "000005ef35fe4002",
		},
		{
			name: "start time is past",
			args: args{
				mainStart: 1593008100,
				increment: 0,
				now:       1593008101,
			},
			want: "000005ef35fe5001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state.mainStart = tt.args.mainStart
			state.increment = tt.args.increment
			if got := getUniqueID(tt.args.now); got != tt.want {
				t.Errorf("genUniqueID() = %v, want %v", got, tt.want)
			}
		})
	}
}
