package protocol

import ()

type Headers map[string]string

type Request struct {
	Path    string
	Method  string
	Headers Headers
}

type Response struct {
	Body    string
	Status  int
	Headers Headers
}

type Expectation struct {
	Request  Request
	Response Response
}

func (r *Request) Cmp(another Request) bool {
	return r.Path == another.Path && r.Method == another.Method
}
