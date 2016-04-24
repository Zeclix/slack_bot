package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

type BusInfo struct {
	mBus    []string
	aBus    []string
	success bool
}

var (
	before_noon []string = []string{"남문", "제2공학관", "과학원", "광복관", "외솔관", "성암관(청송대)", "새천년관", "동문", "경복궁역"}
	after_noon  []string = []string{"남문", "제2공학관", "과학원", "광복관", "외솔관", "성암관(청송대)", "아식설계연구소", "무악학사"}
)

func shuttleErrorMessage(bot *BaseBot, channel string) {
	postResponse(bot, channel, ":bus:", "신촌 셔틀버스", "정보를 가져오는데 에러가 발생했습니다.\nhttp://www.yonsei.ac.kr/sc/campus/traffic1.jsp 에서 직접 확인해주세요.")
}

func getStationFromIndex(arr []string, index int) string {
	arr_len := len(arr)
	if index < arr_len {
		return arr[index]
	} else {
		return arr[arr_len-(index-arr_len+1)-1]
	}
}

func getPositionFromIndex(arr []string, index int) string {
	if index%2 == 0 {
		index /= 2
		return getStationFromIndex(arr, index)
	} else {
		index /= 2
		return fmt.Sprintf("%s -> %s", getStationFromIndex(arr, index), getStationFromIndex(arr, index+1))
	}
}

func getShuttlePosition(field *slack.AttachmentField, arr []string, index int) {
	arr_len := len(arr)

	pos := getPositionFromIndex(arr, index)

	field.Title = pos
	var direction string
	if index >= (arr_len-1)*2 {
		direction = arr[arr_len-1]
	} else {
		direction = arr[0]
	}
	field.Value = fmt.Sprintf("%s 방향", direction)
	field.Short = true
}

func processShuttleCommand(bot *BaseBot, channel string) {
	resp, err := http.Get("http://www.yonsei.ac.kr/_custom/yonsei/_common/shuttle_bus/get_bus_info.jsp")
	if err != nil {
		shuttleErrorMessage(bot, channel)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	b := buf.Bytes()

	var info BusInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		shuttleErrorMessage(bot, channel)
		return
	}

	attachments := make([]slack.AttachmentField, len(info.mBus)+len(info.aBus))
	index := 0
	for pos := range info.mBus {
		getShuttlePosition(&attachments[index], before_noon, pos)
		index++
	}
	for pos := range info.aBus {
		getShuttlePosition(&attachments[index], after_noon, pos)
		index++
	}

	attachment := slack.Attachment{
		Color: "#1766ff",
	}
	if len(attachments) == 0 {
		attachment.Text = "현재 운영중인 셔틀버스 정보 없음"
	} else {
		attachment.Text = "셔틀버스 위치"
		attachment.Fields = attachments
	}

	bot.PostMessage(channel, "", slack.PostMessageParameters{
		AsUser:    false,
		IconEmoji: ":bus:",
		Username:  "신촌 셔틀버스",
		Attachments: []slack.Attachment{
			attachment,
		},
	})
}
