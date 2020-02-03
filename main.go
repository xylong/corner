package main

import (
	_ "corner/routers"
	_ "corner/system"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
