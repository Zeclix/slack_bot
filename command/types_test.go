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

func TestResponseTypeEnum2Str(t *testing.T) {
	cases := []struct {
		enum ResponseTypeEnum
		str  string
	}{
		{in_channel, "in_channel"},
		{deffered_in_channel, "in_channel"},
		{ephemeral, "ephemeral"},
		{123, ""},
	}

	for _, c := range cases {
		got := fmt.Sprintf("%s", c.enum)
		if got != c.str {
			t.Errorf("%d - expected %q, got %q", c.enum, c.str, got)
		}
	}
}

func TestResponseTypeEnum2Json(t *testing.T) {
	cases := []struct {
		enum ResponseTypeEnum
		str  string
		e    bool
	}{
		{in_channel, `"in_channel"`, false},
		{deffered_in_channel, `"in_channel"`, false},
		{ephemeral, `"ephemeral"`, false},
		{123, "", true},
	}

	for _, c := range cases {
		buf, e := json.Marshal(c.enum)
		if (e != nil) != c.e {
			t.Errorf("Error not occured. expected error")
		}
		got := string(buf)
		if got != c.str {
			t.Errorf("%d - expected %q, got %q", c.enum, c.str, got)
		}
	}
}
