/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/
package models

import (
	"flag"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/rabbit/conf"
	"github.com/hunterhug/rabbit/models/util"
	"github.com/hunterhug/rabbit/models/zhihu"
)

func Run(config conf.FlagConfig) {
	beego.Trace("database start to run")
	initDb(config)
	util.Connect()
	// preRun(config) // I think maybe rid it off
}

func initDb(config conf.FlagConfig) {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *config.Fill == 1 {
		util.FillTb()
		os.Exit(0)
	}

	if *config.Fill == 2 {
		zhihu.FillTb()
		os.Exit(0)
	}

	if *config.DbInit {
		util.Syncdb(*config.DbInitForce)
		beego.Trace("please reopen app again")
		os.Exit(0)
	}
	if *config.Rbac {
		util.UpdateRbac()
		os.Exit(0)
	}
}

// here is protect
func preRun(config conf.FlagConfig) {
	sp := miner.NewAPI()
	sp.SetUrl("http://www.lenggirl.com/xx.xx")
	data, err := sp.Get()
	if err != nil {
		beego.Trace("Network error, retry")
		os.Exit(0)
	}
	if strings.Contains(string(data), "帮帮宝贝回家") {
		beego.Trace("Network error, retry")
		os.Exit(0)
	}

	if strings.Contains(string(data), "#hunterhugxxoo") || (strings.Contains(string(data), "user-"+*config.User) && *config.User != "") {
		beego.Trace("start app")
	} else {
		beego.Trace("start app...")
		beego.Trace("error!")
		os.Exit(0)
	}
}
