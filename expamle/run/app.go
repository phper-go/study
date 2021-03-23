package run

import (
	"github.com/phper-go/frame/web/core"
)

var (
	appname = "example"
	App     = &app{}
)

type app struct {
	core.App
}

func init() {
	// core.AppName = appname
}
