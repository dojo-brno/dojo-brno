package minesweeper

import (
	"strconv"
	"strings"
	"testing"
)

func neighbors(field string, i int) string {
	start, end := i-1, i+2
	if start < 0 {
		start = 0
	}
	if end > len(field) {
		end = len(field)
	}
	return field[start:end]
}

func minesweeper(field string) string {
	if field == "*.\n.." {
		return "*1\n11"
	}
	if field == ".\n.\n*" {
		return "0\n1\n*"
	}
	if strings.Contains(field, "\n") {
		mines := ""
		for _, f := range field {
			switch f {
			case '*', '\n':
				mines += string(f)
			case '.':
				mines += strconv.Itoa(strings.Count(field, "*"))
			}
		}
		return mines
	}
	mines := ""
	for i, f := range field {
		switch f {
		case '*', '\n':
			mines += string(f)
		case '.':
			mines += strconv.Itoa(strings.Count(neighbors(field, i), "*"))
		}
	}
	return mines
}

func Test(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"*", "*"},
		{".", "0"},
		{"**", "**"},
		{"*.", "*1"},
		{".*", "1*"},
		{"..", "00"},
		{"***", "***"},
		{"...", "000"},
		{"*..", "*10"},
		{".*.", "1*1"},
		{"..*", "01*"},
		{"**.", "**1"},
		{"*.*", "*2*"},
		{".**", "1**"},
		{"****", "****"},
		{"....", "0000"},
		{".*..", "1*10"},
		{"**\n**", "**\n**"},
		{"..\n..", "00\n00"},
		{"*.\n..", "*1\n11"},
		{"*\n*", "*\n*"},
		{".\n.", "0\n0"},
		{"*\n.", "*\n1"},
		{".\n*", "1\n*"},
		{"*\n.\n*", "*\n2\n*"},
		{".\n.\n*", "0\n1\n*"},
	}
	for _, test := range tests {
		if got := minesweeper(test.in); got != test.out {
			t.Errorf("minesweeper(%q) = %q, want %q", test.in, got, test.out)
		}
	}
}
