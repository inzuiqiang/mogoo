package mogoo

import (
	"fmt"

	"mogoo.me/mogoo/config"
)

var (
	cfg *config.Config
)

func init() {
	cfg_fname := "server.conf"
	conf, err := config.Read(cfg_fname, config.ALTERNATIVE_COMMENT, "=", false, false)
	if nil != err {
		fmt.Errorf("%v\n", err)
		conf = config.NewDefault()
	}
	cfg = conf
}
