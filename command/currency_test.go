package command

import "testing"

func TestCurrencyCommand(t *testing.T) {
	ret := CurrencyCommand(Request{Text: "10 usd = jpy"})
	if ret.ResponseType != deffered_in_channel {
		t.Error("Response error")
	}
	t.Log(ret)
}

func TestSpecialCurrencyCommand(t *testing.T) {
	ret := CurrencyCommand(Request{Text: "10 개리엇 = usd"})
	if ret.ResponseType != deffered_in_channel {
		t.Error("Response error")
	}
	t.Log(ret)
}
