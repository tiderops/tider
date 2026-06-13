package model

import (
	"testing"
	"time"
)

func TestFormatAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{name: "zero time", t: time.Time{}, want: ""},
		{name: "seconds", t: now.Add(-45 * time.Second), want: "45s"},
		{name: "minutes", t: now.Add(-12 * time.Minute), want: "12m"},
		{name: "hours and minutes", t: now.Add(-3*time.Hour - 25*time.Minute), want: "3h25m"},
		{name: "exact hours", t: now.Add(-5 * time.Hour), want: "5h"},
		{name: "days and hours", t: now.Add(-4*24*time.Hour - 6*time.Hour), want: "4d6h"},
		{name: "exact days", t: now.Add(-7 * 24 * time.Hour), want: "7d"},
		{name: "future time clamps to zero", t: now.Add(time.Hour), want: "0s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatAge(tt.t); got != tt.want {
				t.Errorf("FormatAge(%v) = %q, want %q", tt.t, got, tt.want)
			}
		})
	}
}
