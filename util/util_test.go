package util

import (
	"testing"
)

func TestIsEqualStringSlice(t *testing.T) {
	cases := []struct {
		in1, in2 []string
		want     bool
	}{
		{[]string{"Hello,", "world!"}, []string{"Hello,", "world!"}, true},
		{nil, nil, true},
		{[]string{""}, nil, false},
		{nil, []string{"a", "b"}, false},
		{[]string{"aaa", "bbb"}, []string{"aaa"}, false},
		{[]string{"aaa", "bbb"}, []string{"aaa", "bba"}, false},
	}
	for _, c := range cases {
		got := isEqualStringSlice(c.in1, c.in2)
		if got != c.want {
			t.Errorf("isEqualStringSlice(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

func isEqualStringSlice(a, b []string) bool {
	if (a == nil || b == nil) && !(a == nil && b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestDeleteEmpty(t *testing.T) {
	cases := []struct {
		in, want []string
	}{
		{[]string{"", "Hello,", "world!", ""}, []string{"Hello,", "world!"}},
		{[]string{"aaa", "bbb"}, []string{"aaa", "bbb"}},
		{[]string{""}, nil},
		{[]string{"aaa", "", "", ""}, []string{"aaa"}},
	}
	for _, c := range cases {
		got := DeleteEmpty(c.in)
		if !isEqualStringSlice(got, c.want) {
			t.Errorf("DeleteEmpty(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
