package main

import (
	_ "zhiliao_blob/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

