package command

import "testing"

func TestCurrencyCommand(t *testing.T) {
	ret := CurrencyCommand(Request{Text: "10 USD"})
	if ret.ResponseType != deffered_in_channel {
		t.Error("Response error")
	}
	t.Log(ret)
}
