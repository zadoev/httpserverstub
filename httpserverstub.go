package main

import (
	"container/list"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/zadoev/httpserverstub/assertion"
	"github.com/zadoev/httpserverstub/logging"
	"github.com/zadoev/httpserverstub/protocol"
)

const CONTROL_HEADER = "X-Stuby-Control"
const POST = "POST"

const HTTP_405 = http.StatusMethodNotAllowed
const HTTP_400 = http.StatusMethodNotAllowed

const CMD_EXPECT = "expect"
const CMD_DEFAULTS = "defaults"
const CMD_ASSERT = "assert"

var expectations = list.New()

func onError(w http.ResponseWriter, status int, msg string, details string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
	logging.Error.Println(msg, details)
}

func handler(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get(CONTROL_HEADER)

	if action != "" {
		switch action {
		case CMD_EXPECT:
			logging.Trace.Println("Writing expectations")
			if r.Method != POST {
				onError(w, HTTP_405, "Method not allowed", r.Method)
				return
			}
			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				onError(w, HTTP_400, "Can not read request body", "")
				return
			}

			var expectation protocol.Expectation

			err = json.Unmarshal(body, &expectation)

			if err != nil {
				logging.Error.Printf(
					"Can not parse expectation %v with error %v",
					string(body),
					err)
				return
			}

			logging.Trace.Printf("Adding %#v", expectation)

			expectations.PushBack(expectation)

		case CMD_DEFAULTS:

			logging.Trace.Println("Set up defaults")
		case CMD_ASSERT:

			status, message := assertion.Report()

			w.WriteHeader(status)
			io.WriteString(w, message)
		default:

			onError(w, HTTP_400, "Bad request", "Unknown action: "+action)
		}
	} else {
		request := protocol.Request{
			Path:    r.URL.Path,
			Method:  r.Method,
			Headers: protocol.Headers{}}

		var expectation protocol.Expectation

		front_item := expectations.Front()
		if front_item != nil {
			expectation = front_item.Value.(protocol.Expectation)
			expectations.Remove(front_item)
		}

		response := assertion.Assert(&expectation, request)

		w.WriteHeader(response.Status)
		io.WriteString(w, response.Body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	logging.Trace.Println("Starting server at http://127.0.0.1:8181/")
	http.ListenAndServe(":8181", nil)
}
