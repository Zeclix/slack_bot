package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type CommandsInfo struct {
	Port int
}

type CommandInfo map[string]*struct {
	Token   string
	Options []string
}

type CommandRuntimeInfo struct {
	Token   string
	Handler interface{}
	Options map[string]string
}

type CommandServer struct {
	Common   CommandsInfo
	Command  CommandInfo
	Handlers map[string]*CommandRuntimeInfo
}

func NewServer(commands CommandsInfo, command CommandInfo) *CommandServer {
	server := &CommandServer{commands, command, map[string]*CommandRuntimeInfo{}}

	for k, v := range command {
		var parsed_options map[string]string
		for _, val := range v.Options {
			vals := strings.Split(val, ":")
			parsed_options[vals[0]] = vals[1]
		}
		server.Handlers[k] = &CommandRuntimeInfo{v.Token, nil, parsed_options}
	}

	server.registHandler("/echo", EchoCommand, nil)
	server.registHandler("/namu", NamuCommand, nil)
	server.registHandler("/zzal", ZzalCommand, nil)
	server.registHandler("/currency", CurrencyCommand, nil)

	return server
}

type HandlerInitializer func(*map[string]string)

func (server *CommandServer) registHandler(key string, handler interface{}, initializer HandlerInitializer) {
	if val, ok := server.Handlers[key]; ok {
		val.Handler = handler
	} else {
		log.Println("Warning : config not found for ", key)
		server.Handlers[key] = &CommandRuntimeInfo{"", handler, nil}
	}
	if initializer != nil {
		initializer(&server.Handlers[key].Options)
	}
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
				log.Println("Deffered : ", string(buf))
			}

			if e != nil {
				log.Println("Error occured : ", req, e)
			}
		}
	}
}

func (server *CommandServer) Start(wg *sync.WaitGroup) {
	http.HandleFunc("/", server.commandHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", server.Common.Port), nil)

	wg.Done()
}
