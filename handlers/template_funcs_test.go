package handlers

import (
	"testing"

	"time"

	"github.com/google/go-github/github"
)

func TestUTCToLocal(t *testing.T) {
	type args struct {
		t github.Timestamp
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "successful case min",
			args: args{
				t: github.Timestamp{time.Now().Add(time.Minute * 6)},
			},
			want: "5 minutes from now",
		},
		{
			name: "successful case hr",
			args: args{
				t: github.Timestamp{time.Now().Add(time.Hour * 6)},
			},
			want: "5 hours from now",
		},
		{
			name: "successful case d",
			args: args{
				t: github.Timestamp{time.Now().Add(time.Hour * 24 * 6)},
			},
			want: "5 days from now",
		},
		{
			name: "successful case -min",
			args: args{
				t: github.Timestamp{time.Now().Add(-time.Minute * 6)},
			},
			want: "6 minutes ago",
		},
		{
			name: "successful case -hr",
			args: args{
				t: github.Timestamp{time.Now().Add(-time.Hour * 6)},
			},
			want: "6 hours ago",
		},
		{
			name: "successful case -d",
			args: args{
				t: github.Timestamp{time.Now().Add(-time.Hour * 24 * 6)},
			},
			want: "6 days ago",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UTCToLocal(tt.args.t); got != tt.want {
				t.Errorf("UTCToLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadableBoolean(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "True",
			args: args{
				v: true,
			},
			want: "Yes",
		},
		{
			name: "False",
			args: args{
				v: false,
			},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadableBoolean(tt.args.v); got != tt.want {
				t.Errorf("ReadableBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}
