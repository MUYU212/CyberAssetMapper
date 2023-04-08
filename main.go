package main

import (
	"CyberAssetMapper/src/cmd"
	_ "embed"
)

// @title 资产测绘系统 API Doc
// @version 0.0.1
// @description 资产测绘系统
func main() {

	defer cmd.Clean()
	cmd.Start()
	//token, _ := utils.GenerateToken(1, "zs")
	//fmt.Println(token)
	//fmt.Println(utils.IsTokenValid(token))
	//fmt.Println(utils.IsTokenValid(token + "xxxx"))

}
