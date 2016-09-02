package mogoo

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"mogoo.me/mogoo/log"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "", log.Ldefault)
}

type HttpDispatcher struct {
}

func (httpDispatcher HttpDispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := &HttpRequest{Request: r}
	response := &HttpResponse{ResponseWriter: w}
	httpDispatcher.dispatch(request, response)
}

func (httpDispatcher *HttpDispatcher) dispatch(httpRequest *HttpRequest, httpResponse *HttpResponse) {
	ctx := &HttpContext{
		request:  httpRequest,
		response: httpResponse,
	}

	if http.MethodGet == httpRequest.Method() {
		httpDispatcher.doGet(ctx)
	} else if http.MethodPost == httpRequest.Method() {
		httpDispatcher.doPost(ctx)
		fmt.Fprintf(httpResponse.ResponseWriter, "Post : %q\n", ctx)
	} else if http.MethodPut == httpRequest.Method() {
		httpDispatcher.doPut(ctx)
	} else if http.MethodDelete == httpRequest.Method() {
		httpDispatcher.doDelete(ctx)
	} else if http.MethodHead == httpRequest.Method() {
		httpDispatcher.doHead(ctx)
	}
}

func (httpDispatcher *HttpDispatcher) doGet(httpContext *HttpContext) {
	req := httpContext.request

	log := fmt.Sprintf("\"%s %s %s\"", req.Method(), req.URL.Path, req.Proto)
	logger.Infof("%d INFO %s %s\n",
		server.pid,
		NewSimpleDateFormat("[yyyy-MM-dd HH:mm:ss.ffff]").Format(time.Now()),
		log)
}

func (httpDispatcher *HttpDispatcher) doPost(httpContext *HttpContext) {

}

func (httpDispatcher *HttpDispatcher) doPut(httpContext *HttpContext) {

}

func (httpDispatcher *HttpDispatcher) doDelete(httpContext *HttpContext) {

}

func (httpDispatcher *HttpDispatcher) doHead(httpContext *HttpContext) {

}

func (httpDispatcher *HttpDispatcher) doOptions(httpContext *HttpContext) {

}
