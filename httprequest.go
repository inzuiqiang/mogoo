package mogoo

import (
	"net/http"
)

type HttpRequest struct {
	*http.Request
}

func (req *HttpRequest) Method() string {
	return req.Request.Method
}
