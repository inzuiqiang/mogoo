package mogoo

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"mogoo.me/mogoo/config"
)

var (
	server *HttpServer
)

func init() {
	server = newHttpServer()
}

type CommonServer struct {
	server *http.Server
	pid    int
}

func (server *CommonServer) run() error {
	fmt.Printf("%s %d INFO ? mogoo: mogoo is running at %s\n",
		NewSimpleDateFormat("yyyy-MM-dd HH:mm:ss.ffff").Format(time.Now()),
		server.pid,
		server.server.Addr)
	err := server.server.ListenAndServe()
	return err
}

type HttpServer struct {
	*CommonServer
}

func newHttpServer() *HttpServer {
	pid := os.Getpid()

	host, err := cfg.String(config.DEFAULT_SECTION, "host")
	if nil != err {
		host = "127.0.0.1"
	}
	port, err := cfg.Int(config.DEFAULT_SECTION, "port")
	if nil != err {
		port = 8086
	}
	readTimeout, err := cfg.Int(config.DEFAULT_SECTION, "read_timeout")
	if nil != err {
		readTimeout = 60
	}
	writeTimeout, err := cfg.Int(config.DEFAULT_SECTION, "write_timeout")
	if nil != err {
		writeTimeout = 60
	}
	addr := fmt.Sprintf("%s:%d", host, port)

	return &HttpServer{
		CommonServer: &CommonServer{
			server: &http.Server{
				Addr:           addr,
				Handler:        HttpDispatcher{},
				ReadTimeout:    time.Duration(readTimeout) * time.Second,
				WriteTimeout:   time.Duration(writeTimeout) * time.Second,
				MaxHeaderBytes: 1 << 20,
			},
			pid: pid,
		},
	}
}
