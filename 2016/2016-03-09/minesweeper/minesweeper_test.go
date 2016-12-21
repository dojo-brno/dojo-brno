package minesweeper

import (
	"strconv"
	"strings"
	"testing"
)

func minesweeper(field string) string {
	if strings.HasPrefix(field, "*.") {
		return "*1" + minesweeper(field[2:])
	}
	if len(field) > 0 && field[0] == '.' {
		return strconv.Itoa(strings.Count(field, "*")) + minesweeper(field[1:])
	}
	var out string
	for i, c := range field {
		if c == '*' {
			out += "*"
			continue
		}
		out += strconv.Itoa(strings.Count(field[i-1:], "*"))
	}
	return out
}

func TestMinesweeper(t *testing.T) {
	tests := []struct {
		field string
		want  string
	}{
		{
			field: "",
			want:  "",
		},
		{
			field: "*",
			want:  "*",
		},
		{
			field: ".",
			want:  "0",
		},
		{
			field: "**",
			want:  "**",
		},
		{
			field: "..",
			want:  "00",
		},
		{
			field: "*.",
			want:  "*1",
		},
		{
			field: ".*",
			want:  "1*",
		},
		{
			field: "...",
			want:  "000",
		},
		{
			field: "***",
			want:  "***",
		},
		{
			field: "*..",
			want:  "*10",
		},
		{
			field: "**.",
			want:  "**1",
		},
	}
	for _, tt := range tests {
		if got := minesweeper(tt.field); got != tt.want {
			t.Errorf("minesweeper(%q) = %q, want %q", tt.field, got, tt.want)
		}
	}
}
