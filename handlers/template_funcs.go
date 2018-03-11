package handlers

import (
	"github.com/dustin/go-humanize"
	"github.com/google/go-github/github"
)

// UTCToLocal used as template helper to format UTC time to local with
// human readable format.
func UTCToLocal(t github.Timestamp) string {
	return humanize.Time(t.Local())
}

// ReadableBoolean used as template helper to format boolean into human readable format.
func ReadableBoolean(v bool) string {
	if v {
		return "Yes"
	}
	return "No"
}
