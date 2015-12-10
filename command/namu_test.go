package command

import "testing"

func TestNamuCommandManual(t *testing.T) {
	var req Request
	req.Text = ""

	res := NamuCommand(req)
	if res.ResponseType != ephemeral {
		t.Errorf("Response type mismatch, wanted ephemeral but %s", res.ResponseType)
	}
	if len(res.Attachments) == 0 {
		t.Error("Attachments not found")
	}
	att := res.Attachments[0]
	if att.Pretext == "" {
		t.Error("Pretext is empty")
	}
}

func TestNamuCommand(t *testing.T) {
	cases := []struct {
		text, username       string
		responseType         ResponseTypeEnum
		title, link, resText string
	}{
		{"피카츄", "s_y_e_2___",
			deffered_in_channel,
			"나무위키 - 피카츄",
			"https://namu.wiki/w/%ED%94%BC%EC%B9%B4%EC%B8%84",
			"@s_y_e_2___ 나무위키 링크"},
		{"몰라 뭐야 그거 무서워", "s_y_e_2___",
			deffered_in_channel,
			"나무위키 - 몰라 뭐야 그거 무서워",
			"https://namu.wiki/w/%EB%AA%B0%EB%9D%BC%20%EB%AD%90%EC%95%BC%20%EA%B7%B8%EA%B1%B0%20%EB%AC%B4%EC%84%9C%EC%9B%8C",
			"@s_y_e_2___ 나무위키 링크"},
		{" remove useless spaces  ", "s_y_e_2___",
			deffered_in_channel,
			"나무위키 - remove useless spaces",
			"https://namu.wiki/w/remove%20useless%20spaces",
			"@s_y_e_2___ 나무위키 링크"},
	}
	for _, c := range cases {
		var req Request
		req.Text = c.text
		req.UserName = c.username

		res := NamuCommand(req)
		if len(res.Attachments) == 0 {
			t.Errorf("Error attachments not found")
		}
		att := res.Attachments[0]
		if res.ResponseType != c.responseType {
			t.Errorf("ResponseType mismatch expected %s, but %s", c.responseType, res.ResponseType)
		}
		if att.Title != c.title {
			t.Errorf("Title mismatch expected %s, but %s", c.title, att.Title)
		}
		if att.TitleLink != c.link {
			t.Errorf("TitleLink mismatch expected %s, but %s", c.link, att.TitleLink)
		}
		if att.Text != c.resText {
			t.Errorf("Text mismatch expected %s, but %s", c.resText, att.Text)
		}
	}
}
