package command

import (
	"strings"
	"testing"
)

func TestZzalCommandSlugSuccess(t *testing.T) {
	var req Request
	req.Text = "no-hope-no-dream"
	req.UserName = "user_name"

	var respType ResponseTypeEnum
	respType = deffered_in_channel
	imgUrl := "http://zzal.collapsed.me/posts/no-hope-no-dream/OTL.jpg"

	res := ZzalCommand(req)
	if res.ResponseType != respType {
		t.Errorf("ResponseType Error : expected %q, but %q", respType, res.ResponseType)
	}
	if len(res.Attachments) == 0 {
		t.Error("Attachments not found")
	}
	att := res.Attachments[0]
	if att.ImageUrl != imgUrl {
		t.Errorf("Image URL mismatch : expected %q, but %q", imgUrl, att.ImageUrl)
	}
}

func TestZzalCommandSlugError(t *testing.T) {
	var req Request
	req.Text = "yes-hope-yes-dream"
	req.UserName = "user_name"

	var respType ResponseTypeEnum
	respType = ephemeral

	res := ZzalCommand(req)
	if res.ResponseType != respType {
		t.Errorf("ResponseType Error : expected %q, but %q", respType, res.ResponseType)
	}
	if len(res.Attachments) != 0 {
		t.Error("Error : Attachments found")
	}
	if !strings.HasPrefix(res.Text, "에러:") {
		t.Errorf("Error : no error report (%q)", res.Text)
	}
}

func TestZzalCommandManual(t *testing.T) {
	var req Request

	res := ZzalCommand(req)
	if res.ResponseType != ephemeral {
		t.Errorf("ResponseType mismatch : expedted %q, but %q", ephemeral, res.ResponseType)
	}
	if len(res.Attachments) == 0 {
		t.Errorf("Attachments not found")
	}
}
