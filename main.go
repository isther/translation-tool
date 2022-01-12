package main

import (
	"fmt"

	"github.com/isther/translation-tool/cmd"
	"github.com/isther/translation-tool/config"
	"github.com/isther/translation-tool/util"
)

const configPath = "/.translation/config"

// 接口使用有道智云
// https://ai.youdao.com/
func init() {
	home, _ := util.Home()
	cfgPath := fmt.Sprintf("%v%v", home, configPath)
	config.Init(cfgPath)
}

func main() {
	cmd.Execute()
}
