package admin

import (
	"github.com/phper-go/frame/web/core"
)

type Index struct {
	core.Controller
}

func (this *Index) IndexAction() {
	this.Output().Content = []byte("/index/index ok")
}

func init() {
	core.RegisterController("/index", &Index{})
}
