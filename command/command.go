package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"net/http"
	"reflect"
	"sync"
)

type CommandsInfo struct {
	Port int
}

type CommandInfo map[string]*struct {
	Token string
}

type CommandRuntimeInfo struct {
	Token   string
	Handler interface{}
}

type CommandServer struct {
	Common   CommandsInfo
	Command  CommandInfo
	Handlers map[string]*CommandRuntimeInfo
}

type AttachmentField struct {
	title string
	value string
	short bool
}

type Attachment struct {
	fallback string
	color    color.Color
	pretext  string

	author_name string
	author_link string
	author_icon string

	title      string
	title_link string

	text string

	image_url string
	thumb_url string
}

type ResponseTypeEnum int

const (
	in_channel = iota
	ephemeral
	deffered_in_channel
)

func (e ResponseTypeEnum) MarshalJSON() ([]byte, error) {
	var str string
	switch e {
	case deffered_in_channel:
	case in_channel:
		str = "in_channel"
		break
	case ephemeral:
		str = "ephemeral"
		break
	default:
		return nil, errors.New("Invalid value")
	}
	return []byte(`"` + str + `"`), nil
}

type Response struct {
	ResponseType ResponseTypeEnum `json:"response_type"`
	Text         string           `json:"text"`
	Attachments  []Attachment     `json:"attachments"`
}

type Request struct {
	Token       string `param:"token"`
	TeamId      string `param:"team_id"`
	TeamDomain  string `param:"team_domain"`
	ChannelId   string `param:"channel_id"`
	ChannelName string `param:"channel_name"`
	UserId      string `param:"user_id"`
	UserName    string `param:"user_name"`
	Command     string `param:"command"`
	Text        string `param:"text"`
	ResponseUrl string `param:"response_url"`
}

func NewServer(commands CommandsInfo, command CommandInfo) *CommandServer {
	server := &CommandServer{commands, command, map[string]*CommandRuntimeInfo{}}

	for k, v := range command {
		server.Handlers[k] = &CommandRuntimeInfo{v.Token, nil}
	}

	server.Handlers["/echo"].Handler = EchoCommand

	return server
}

func requestFormToRequestObj(r *http.Request) *Request {
	ret := new(Request)

	val := reflect.Indirect(reflect.ValueOf(ret))
	typ := reflect.TypeOf(*ret)

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		field_info := typ.Field(i)
		field_name := field_info.Tag.Get("param")
		field.Set(reflect.ValueOf(r.FormValue(field_name)))
	}

	return ret
}

func (server *CommandServer) commandHandler(w http.ResponseWriter, r *http.Request) {
	req := requestFormToRequestObj(r)
	handlerInfo := server.Handlers[req.Command]

	if handlerInfo != nil {
		if handlerInfo.Token == "" || handlerInfo.Token == req.Token {
			fun := reflect.ValueOf(handlerInfo.Handler)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(*req)
			response := fun.Call(in)[0].Interface().(*Response)

			var e error
			w.Header().Set("Content-Type", "application/json")
			if response.ResponseType != deffered_in_channel {
				encoder := json.NewEncoder(w)
				e = encoder.Encode(response)
			} else {
				var buf []byte
				buf, e = json.Marshal(response)
				http.Post(req.ResponseUrl, "application/json", bytes.NewBuffer(buf))
			}

			if e != nil {
				fmt.Println("Error occured : ", req, e)
			}
		}
	}
}

func (server *CommandServer) Start(wg *sync.WaitGroup) {
	http.HandleFunc("/", server.commandHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", server.Common.Port), nil)

	wg.Done()
}
