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

func TestCurrencyAlias(t *testing.T) {
	unit := "사대강"
	applyAlias(&unit)
	if unit != "4대강" {
		t.Error("alias error")
	}
}
