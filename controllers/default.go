package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.Data["Hello"]="Hello world"
	//c.TplName = "index.tpl"
	//c.Ctx.WriteString("hahahhah")
	fmt.Println("XXX")
	type JJ struct {
	Name  string
	Name1 string
}
	//`json:"name"`
	js:=JJ{Name: "n55ame"}
	js.Name ="xxx"
	js.Name1 ="oooo"
	//c.Data["json"]=&js
	c.Data["json"]=js
	c.ServeJSON()
}
