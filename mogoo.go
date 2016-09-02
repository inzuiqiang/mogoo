package mogoo

import "mogoo.me/mogoo/text"

func NewSimpleDateFormat(format string) *text.SimpleDateFormat {
	return text.NewSimpleDateFormat(format)
}

func Run() error {
	return server.run()
}
