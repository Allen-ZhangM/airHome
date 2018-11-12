package main

import (
	_ "airHome/routers"
	"github.com/astaxie/beego"
	_ "airHome/models"
	"airHome/blockchain"
	"os"
	"fmt"
)

func main() {
	//定义SDK属性
	fSetup:=blockchain.FabricSetup{
		OrgAdmin:"Admin",
		OrgName:"Org1",
		ConfigFile:"config.yaml",
		//通道相关
		ChannelID:"mychannel",
		ChannelConfig:os.Getenv("GOPATH")+"/src/airHome/fixtures/artifacts/channel.tx",
	}
	//初始化SDK
	err:=fSetup.Initialize()
	if err!=nil{
		fmt.Printf("无法初始化Fabric SDK: %v\n", err)
	}
	beego.Run()
}



