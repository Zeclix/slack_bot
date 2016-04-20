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
	command_re *regexp.Regexp       = regexp.MustCompile("(\\d+(?:\\.\\d+)?)\\s*([^\\s=]+)(\\s*=\\s*[?]?\\s*([^\\s=]+))?")
	rate_cache map[string]rateCache = map[string]rateCache{}
	alias_map map[string]string = map[string]string {
		"사대강": "4대강",
		"K달탐사": "k달탐사",
		"케이달탐사": "k달탐사",
		"K알파고": "k알파고",
		"케이알파고": "k알파고",
		"tmax해킹": "티맥스해킹",
		"TMAX해킹": "티맥스해킹",
	}
	special_currency_map map[string]currencyRatio = map[string]currencyRatio{
		"개리엇": currencyRatio{ 30000000, "USD" },
		"4대강" : currencyRatio{ 22000000000000, "KRW" },
		"팝픽" : currencyRatio{ 344, "KRW" },
		"k달탐사" : currencyRatio{ 10000000000, "KRW" },
		"k알파고" : currencyRatio{ 30000000000, "KRW" },
		"최저시급": currencyRatio{ 6030, "KRW" },
		"티맥스해킹": currencyRatio{ 100000000, "KRW" },
	}
)

func applyAlias(unit *string) {
	if alias, ok := alias_map[*unit]; ok {
		*unit = alias
	}
}

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

	src_value := original_value
	// make upper...
	matched[2] = strings.ToUpper(matched[2])
	applyAlias(&matched[2])
	matched[4] = strings.ToUpper(matched[4])
	applyAlias(&matched[4])

	src_unit := matched[2]
	target_unit := matched[4]
	if info, ok := special_currency_map[src_unit]; ok {
		src_value = src_value * info.rate
		src_unit = info.unit
	}
	if info, ok := special_currency_map[target_unit]; ok {
		src_value = src_value / info.rate
		target_unit = info.unit
	}
	key := fmt.Sprintf("%s%s", src_unit, target_unit)

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
			Text:  fmt.Sprintf("%.2f %s = %.2f %s", original_value, matched[2], src_value*rate, matched[4]),
		},
	}

	return ret
}
