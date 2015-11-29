package command

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestColor2String(t *testing.T) {
	cases := []struct {
		color Color
		str   string
	}{
		{Color{0, 0, 0}, "#000000"},
		{Color{255, 0, 0}, "#ff0000"},
		{Color{0, 128, 0}, "#008000"},
		{Color{0, 0, 15}, "#00000f"},
	}

	for _, c := range cases {
		got := fmt.Sprintf("%s", c.color)
		if got != c.str {
			t.Errorf("Color{%d, %d, %d} - expected %q, got %q", c.color.r, c.color.g, c.color.b, c.str, got)
		}
	}
}

func TestColor2Json(t *testing.T) {
	cases := []struct {
		color Color
		str   string
	}{
		{Color{0, 0, 0}, `"#000000"`},
		{Color{255, 0, 0}, `"#ff0000"`},
		{Color{0, 128, 0}, `"#008000"`},
		{Color{0, 0, 15}, `"#00000f"`},
	}

	for _, c := range cases {
		buf, e := json.Marshal(c.color)
		if e != nil {
			t.Errorf("Error occured while converting to json : %q", e)
		}
		got := string(buf)
		if got != c.str {
			t.Errorf("json.Marshal(Color{%d, %d, %d}) - expected %q, got %q", c.color.r, c.color.g, c.color.b, c.str, got)
		}
	}
}
