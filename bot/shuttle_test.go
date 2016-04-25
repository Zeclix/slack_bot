package bot

import "testing"

func TestPositionToString(t *testing.T) {
	cases := []struct {
		input    int
		target   []string
		expected string
	}{
		{0, before_noon, "남문"},
		{1, before_noon, "남문 -> 제2공학관"},
		{15, before_noon, "동문 -> 경복궁역"},
		{16, before_noon, "경복궁역"},
		{18, before_noon, "동문"},
		{17, before_noon, "경복궁역 -> 동문"},
		{32, before_noon, "남문"},
		{0, after_noon, "남문"},
		{14, after_noon, "무악학사"},
		{13, after_noon, "아식설계연구소 -> 무악학사"},
	}

	for i, test_case := range cases {
		output := getPositionFromIndex(test_case.target, test_case.input)
		if output != test_case.expected {
			t.Errorf("Failed test for %dth case, Expected %q but %q", i, test_case.expected, output)
		}
	}
}
