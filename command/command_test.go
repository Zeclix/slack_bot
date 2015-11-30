package command

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRequestFormToRequstObj(t *testing.T) {
	req := &http.Request{Method: "GET"}
	req.URL, _ = url.Parse("http://test.com/?token=token&channel_id=channel_id&ChannelId=ChannelId")
	res := requestFormToRequestObj(req)
	if res.Token != "token" {
		t.Error("wrong value is set.")
	}
	if res.ChannelId == "ChannelId" {
		t.Error("wrong value is set. maybe tag is wrong...")
	}
	if res.ChannelId != "channel_id" {
		t.Error("wrong value is set.")
	}
}
