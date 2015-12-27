package command

import "testing"

func TestCurrencyCommand(t *testing.T) {
	ret := CurrencyCommand(Request{Text: "1USD=?KRW"})
	if ret.ResponseType != deffered_in_channel {
		t.Error("Response error")
	}
}
