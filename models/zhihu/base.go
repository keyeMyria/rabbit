package zhihu

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/rabbit/models/util"
)

func FillTb() {
	beego.Trace("data zhihu fill init start")
	// TruncateRbacTable([]string{"category", "paper", "roll"})
	util.Connect()
	UpdateConfig()
	InsertCategory()
	InsertPaper()
	InsertRoll()
	beego.Trace("data zhihu fill init end")
}
