package command

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CurrencyQuery struct {
	XMLName xml.Name `xml:"query"`
	Rates   []struct {
		ID   string  `xml:"id,attr"`
		Rate float64 `xml:"Rate"`
	} `xml:"results>rate"`
}

type rateCache struct {
	rate      float64
	cached_at time.Time
}

type currencyRatio struct {
	rate float64
	unit string
}

var (
	command_re *regexp.Regexp       = regexp.MustCompile("(\\d+(?:\\.\\d+)?)\\s*(.{3})(\\s*=\\s*[?]?\\s*(.{3}))?")
	rate_cache map[string]rateCache = map[string]rateCache{}
	special_currency_map map[string]currencyRatio = map[string]currencyRatio{
		"개리엇": currencyRatio{ 30000000, "USD" },
		"4대강" : currencyRatio{ 22000000000000, "KRW" },
	}
)

func CurrencyCommand(req Request) *Response {
	ret := new(Response)
	ret.ResponseType = ephemeral

	matched := command_re.FindStringSubmatch(strings.TrimSpace(req.Text))
	if matched == nil {
		ret.Text = fmt.Sprintf("Format error. You typed \"%s\"", req.Text)
		return ret
	}

	original_value, _ := strconv.ParseFloat(matched[1], 64)
	if matched[3] == "" {
		matched[4] = "KRW"
	}
	// make upper...
	matched[2] = strings.ToUpper(matched[2])
	matched[4] = strings.ToUpper(matched[4])
	if info, ok := special_currency_map[matched[2]]; ok {
		original_value = original_value * info.rate
		matched[2] = info.unit
	}
	if info, ok := special_currency_map[matched[4]]; ok {
		original_value = original_value / info.rate
		matched[4] = info.unit
	}
	key := fmt.Sprintf("%s%s", matched[2], matched[4])

	rate := 0.0
	if cached, ok := rate_cache[key]; ok {
		if time.Now().Sub(cached.cached_at).Minutes() < 10 {
			rate = cached.rate
		}
	}

	if rate == 0.0 {
		url := fmt.Sprintf("http://query.yahooapis.com/v1/public/yql?q=%s&env=%s",
			strings.Replace(
				url.QueryEscape(
					fmt.Sprintf("select * from yahoo.finance.xchange where pair in (\"%s\")", key)),
				"+", "%20", -1),
			url.QueryEscape("store://datatables.org/alltableswithkeys"))

		resp, err := http.Get(url)
		if err != nil {
			ret.Text = fmt.Sprintf("YQL error : %q", err)
			return ret
		}

		var query CurrencyQuery
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		b := buf.Bytes()
		err = xml.Unmarshal(b, &query)
		if err != nil {
			ret.Text = fmt.Sprintf("Result parsing error : %q", err)
			return ret
		}

		for _, v := range query.Rates {
			if v.ID == key {
				rate = v.Rate
				break
			}
		}

		if rate != 0.0 {
			rate_cache[key] = rateCache{
				rate:      rate,
				cached_at: time.Now(),
			}
		} else {
			ret.Text = "Unknown Error"
			return ret
		}
	}

	ret.ResponseType = deffered_in_channel

	ret.Attachments = []Attachment{
		Attachment{
			Color: Color{r: 33, g: 108, b: 42},
			Text:  fmt.Sprintf("%.2f %s = %.2f %s", original_value, matched[2], original_value*rate, matched[4]),
		},
	}

	return ret
}
