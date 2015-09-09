package assertion

import (
	"../protocol"
	"container/list"
	"fmt"
	"net/http"
)

type Assertion struct {
	IsOk bool
	Msg  string
}

var assertions = list.New()

func Assert(e *protocol.Expectation, req protocol.Request) protocol.Response {
	var expected_req protocol.Request
	var assertion Assertion
	var response protocol.Response

	if e != nil {
		expected_req = e.Request
	}

	if e != nil && expected_req.Cmp(req) {
		response = e.Response
		assertion = Assertion{IsOk: true, Msg: ""}
	} else {
		// @todo: here we should use defaults
		msg := fmt.Sprintf("No match for: %v %v", req.Method, req.Path)
		assertion = Assertion{IsOk: false, Msg: msg}
		response = protocol.Response{
			Body:    "",
			Status:  http.StatusNotFound,
			Headers: protocol.Headers{}}
	}

	assertions.PushBack(assertion)
	return response
}

func Report() (int, string) {
	isOk := true
	message := ""

	for e := assertions.Front(); e != nil; e = e.Next() {

		isOk = isOk && e.Value.(Assertion).IsOk

		message += "\n" + e.Value.(Assertion).Msg
	}

	var status int

	if isOk {
		status = http.StatusOK
	} else {
		status = http.StatusInternalServerError
	}

	return status, message
}
