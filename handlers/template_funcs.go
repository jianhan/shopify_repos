package handlers

import (
	"github.com/dustin/go-humanize"
	"github.com/google/go-github/github"
)

func UTCToLocal(t github.Timestamp) string {
	return humanize.Time(t.Local())
}

func ReadableBoolean(v bool) string {
	if v {
		return "Yes"
	}
	return "No"
}
