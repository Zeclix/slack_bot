package bot

import (
	. "github.com/PoolC/slack_bot/util"
	"testing"
)

func newAnzu() *Anzu {
	return &Anzu{nil, NewRedisMock()}
}

func TestSimpleReactionAnzu(t *testing.T) {
	anzu := newAnzu()

	cases := []struct {
		text     string
		username string
		want     interface{}
	}{
		{"안즈 카와이", "", "뭐... 뭐라는거야"},
		{"안즈쨩 카와이", "", "뭐... 뭐라는거야"},
		{"안즈카와이", "", nil},
		{"안즈쨩 뭐해?", "", "숨셔"},
		{"사람은 일을 하고 살아야한다. 메우", "meu", "이거 놔라 이 퇴근도 못하는 놈이"},
	}

	for _, c := range cases {
		ret := anzuMessageProcess(anzu, newMessageEvent(c.text, c.username))
		if ret != c.want {
			t.Errorf("Simple reaction test failed. expected \"%q\" but \"%q\"", c.want, ret)
		}
	}
}
