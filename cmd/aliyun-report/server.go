package main

import (
	"github.com/shynome/aliyun-monitor-report/api"
	_ "github.com/shynome/aliyun-monitor-report/api/init"
)

func main() {
	api.Server.Logger.Fatal(api.Server.Start(":3000"))
}
