package command

import "testing"

func TestEchoCommand(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"asdf"},
		{""},
		{"lorem ipsum"},
		{"가나다라 하하핳"},
	}

	for _, c := range cases {
		var req Request
		req.Text = c.in
		got := EchoCommand(req)
		if got.ResponseType != deffered_in_channel {
			t.Error("??!! wrong response type")
		}
		if got.Text != c.in {
			t.Error("??!! wrong text...")
		}
	}
}
