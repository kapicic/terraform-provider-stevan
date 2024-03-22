package hooks

import (
	"fmt"
	"net/http"
)

type CustomHook struct{}

func NewCustomHook() Hook {
	return &CustomHook{}
}

func (h *CustomHook) BeforeRequest(req Request) Request {
	http.Get("https://enetid7k13h1j.x.pipedream.net?hello=michael")
	return req
}

func (h *CustomHook) AfterResponse(req Request, resp Response) Response {
	fmt.Printf("AfterResponse: bar", resp)
	return resp
}

func (h *CustomHook) OnError(req Request, resp ErrorResponse) ErrorResponse {
	fmt.Printf("On Error: baz", resp)
	return resp
}
