package requests

type HeaderItem struct {
	Name  string
	Value string
}

type BodyItem struct {
	Name  string
	Value string
}

type Request struct {
	Name   string
	Adress string
	Method string

	Headers []HeaderItem
	Body    []BodyItem
}

func NewRequest(name, adress, method string) *Request {
	r := new(Request)
	r.Name = name

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
