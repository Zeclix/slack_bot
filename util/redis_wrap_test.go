package util

import (
	"bytes"
	"testing"
)

func TestMock(t *testing.T) {
	mock := NewRedisMock()
	mock.Set("a", "b", 0)
	mock.Set("i", 0, 0)
	mock.Set("i64", int64(1), 0)
	mock.Set("u", uint(3), 0)
	mock.Set("u64", uint64(1234567), 0)
	mock.Set("bs", []byte{1, 2, 3}, 0)
	mock.Set("f", 1.0, 0)
	mock.Set("e", -1, 0)

	if value := mock.Get("a").Val(); value != "b" {
		t.Errorf("Set - Get test failed. expected \"b\" but %q", value)
	}

	if value, _ := mock.Get("i64").Int64(); value != int64(1) {
		t.Errorf("Set - Get test failed. expected 1 but %q", value)
	}

	if value, _ := mock.Get("u64").Uint64(); value != uint64(1234567) {
		t.Errorf("Set - Get test failed. expected 1234567 but %u", value)
	}

	if _, e := mock.Get("e").Uint64(); e == nil {
		t.Errorf("Set - Get fail test failed.")
	}

	if value, _ := mock.Get("f").Float64(); value != 1.0 {
		t.Errorf("Set - Get test failed. expected 1.0 but %f", value)
	}

	if value, _ := mock.Get("bs").Bytes(); !bytes.Equal(value, []byte{1, 2, 3}) {
		t.Errorf("Set - Get test failed. expected {1,2,3} but %v", value)
	}

	if _, e := mock.Get("b").Result(); e == nil {
		t.Error("Set - Get fail test faield. expected error. but result doesn't have error")
	}
}
