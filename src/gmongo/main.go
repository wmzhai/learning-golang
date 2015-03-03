package main
 
import (
	_ "gmongo/docs"
	_ "gmongo/routers"
	"github.com/astaxie/beego"
)

func main() {
//	if beego.RunMode == "dev" {

//	}
	beego.DirectoryIndex = true
	beego.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
