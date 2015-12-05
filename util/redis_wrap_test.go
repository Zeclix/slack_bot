package util

import "testing"

func TestMock(t *testing.T) {
	mock := NewRedisMock()
	mock.Set("a", "b", 0)

	if value := mock.Get("a").Val(); value != "b" {
		t.Errorf("Set - Get test failed. expected \"b\" but %q", value)
	}

	if _, e := mock.Get("b").Result(); e == nil {
		t.Error("Set - Get fail test faield. expected error. but result doesn't have error")
	}
}
