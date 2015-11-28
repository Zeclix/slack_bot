package command

import (
	"fmt"
	"github.com/PoolC/slack_bot/util"
	"golang.org/x/net/html"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var link_re *regexp.Regexp

func init() {
	link_re = regexp.MustCompile("/([^/]+)/$")
}

func ZzalCommand(req Request) *Response {
	tokens := util.DeleteEmpty(strings.Split(req.Text, " "))
	ret := new(Response)

	if len(tokens) == 0 {
		ret.ResponseType = ephemeral
		ret.Text = "사용법"
		var default_usage, random_usage Attachment
		default_usage.Title = "가져오기"
		default_usage.Text = "/zzal [name]"
		random_usage.Title = "랜덤 짤"
		random_usage.Text = `/zzal !random
/zzal !r`
		ret.Attachments = append(ret.Attachments, default_usage, random_usage)
		return ret
	}

	key := tokens[0]
	switch key {
	case "!random":
		fallthrough
	case "!r":
		url := "https://shipduck.github.io/umi/archives.html"
		resp, err := http.Get(url)
		if err != nil {
			ret.ResponseType = ephemeral
			ret.Text = fmt.Sprintf("에러: 목록 가져오기 에러.", key)
			return ret
		}

		doc, err := html.Parse(resp.Body)
		var list []string
		var f func(*html.Node) bool
		f = func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				switch {
				case n.Data == "dd":
					for _, attr := range n.FirstChild.Attr {
						if attr.Key == "href" {
							matched := link_re.FindStringSubmatch(attr.Val)
							list = append(list, matched[1])
							break
						}
					}
					break
				case n.Data == "footer":
					return false
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if !f(c) {
					return false
				}
			}

			return true
		}
		f(doc)

		key = list[rand.Intn(len(list))]
		fallthrough
	default:
		var image Attachment
		post_url := fmt.Sprintf("https://shipduck.github.io/umi/posts/%s/", key)
		image.TitleLink = post_url
		resp, err := http.Get(post_url)
		if err != nil {
			ret.ResponseType = ephemeral
			ret.Text = fmt.Sprintf("에러: %s 짤이 없습니다.", key)
			return ret
		}

		doc, err := html.Parse(resp.Body)
		var f func(*html.Node) bool
		f = func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				switch {
				case n.Data == "meta":
					meta := map[string]string{}
					for _, attr := range n.Attr {
						meta[attr.Key] = attr.Val
					}

					if meta["name"] == "twitter:image:src" {
						parsed_url, _ := url.Parse(meta["content"])
						image.ImageUrl = parsed_url.String()
						log.Println(image)
					}
					break
				case n.Data == "title":
					image.Title = n.FirstChild.Data
					break
				case n.Data == "body":
					return false
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if !f(c) {
					return false
				}
			}

			return true
		}
		f(doc)

		if len(image.ImageUrl) == 0 {
			ret.ResponseType = ephemeral
			ret.Text = "에러: 이미지 정보가 없음"
			return ret
		}

		image.Text = fmt.Sprintf("@%s", req.UserName)
		ret.ResponseType = deffered_in_channel

		ret.Attachments = append(ret.Attachments, image)
	}
	return ret
}
