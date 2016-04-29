package bot

import "testing"

func TestSimpleMatch(t *testing.T) {
	cases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"큿", []string{"큿"}, true},
		{"큿", []string{"큿", "72"}, true},
		{"큿큿", []string{"큿"}, false},
		{"치하야72", []string{"72", "치하야"}, false},
		{"72 치하야 72", []string{"치하야"}, true},
		{"끄아아72", []string{"72"}, false},
	}

	for i, test_case := range cases {
		output := simpleMatch(test_case.input, test_case.patterns...)
		if output != test_case.expected {
			t.Errorf("Failed test for %dth case, Expected %t but %t", i, test_case.expected, output)
		}
	}
}
