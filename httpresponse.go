package mogoo

import (
	"net/http"
)

type HttpResponse struct {
	http.ResponseWriter
}
