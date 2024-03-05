package requests

import (
	"errors"
	"fmt"
	"net/http"
)

type HeaderItem struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type BodyItem struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Request struct {
	Name     string       `json:"Name"`
	Protocol string       `json:"Protocol"`
	Adress   string       `json:"Adress"`
	Method   string       `json:"Method"`
	Headers  []HeaderItem `json:"Headers"`
	Body     []BodyItem   `json:"Body"`
}

func NewRequest(name, adress, method string) *Request {
	r := new(Request)
	r.Name = name
	r.Adress = adress
	r.Method = method
	return r
}

func NewBodyItem(name, value string) *BodyItem {
	item := new(BodyItem)
	item.Name = name
	item.Value = value
	return item
}

func NewHeaderItem(name, value string) *HeaderItem {
	item := new(HeaderItem)
	item.Name = name
	item.Value = value
	return item
}

func (v *Request) ProceedRequest() (*http.Response, error) {
	var res *http.Response = nil
	var err error = nil
	switch v.Method {
	case "get":
		res, err = http.Get(v.Adress)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
		}
	default:
		err = errors.New("Incorrect ")
	}

	return res, err
}
