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

// Command server info
//
// use with gcfg
type CommandsInfo struct {
	Port int
}

// Each command processor info
//
// user with gcfg
type CommandInfo map[string]*struct {
	Token   string
	// "key=val" list
	Options []string
}

// Command processor info using at runtime
type CommandRuntimeInfo struct {
	Token   string
	Handler interface{}
	Options map[string]string
}

// Command server object
type CommandServer struct {
	Common   CommandsInfo
	Command  CommandInfo
	Handlers map[string]*CommandRuntimeInfo
}

// Create new server
func NewServer(commands CommandsInfo, command CommandInfo) *CommandServer {
	server := &CommandServer{commands, command, map[string]*CommandRuntimeInfo{}}

	// set config from config file.
	for k, v := range command {
		var parsed_options map[string]string
		for _, val := range v.Options {
			vals := strings.Split(val, ":")
			parsed_options[vals[0]] = vals[1]
		}
		server.Handlers[k] = &CommandRuntimeInfo{v.Token, nil, parsed_options}
	}

	// Regist processor
	//
	// Reflection of golang can't find function with name.
	// manually regist all available handler here.
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

// Comman request handler
func (server *CommandServer) commandHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request
	req := requestFormToRequestObj(r)
	// retrieve command handler for request
	handlerInfo := server.Handlers[req.Command]

	// exist handler
	if handlerInfo != nil {
		// token check
		if handlerInfo.Token == "" || handlerInfo.Token == req.Token {
			// invoke handler
			fun := reflect.ValueOf(handlerInfo.Handler)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(*req)
			response := fun.Call(in)[0].Interface().(*Response)

			// create response
			var e error
			w.Header().Set("Content-Type", "application/json")
			if response.ResponseType != deffered_in_channel {
				encoder := json.NewEncoder(w)
				e = encoder.Encode(response)
			} else {
				// special case - deffered_in_channel
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

// start server
func (server *CommandServer) Start(wg *sync.WaitGroup) {
	http.HandleFunc("/", server.commandHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", server.Common.Port), nil)

	wg.Done()
}
