package controllers

import (
	"controllers/statistic"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	//fmt.Println(c.Ctx.Request.Body)
	username := c.GetString("username")
	fmt.Println(username)
	c.Data["json"] = map[string]interface{}{"code":20000, "data":tokens["admin"]}
	c.ServeJSON()
}

type GetInfoController struct {
	beego.Controller
}

func (c *GetInfoController) Get() {
	token := c.GetString("token")
	//fmt.Println(token)
	c.Data["json"] = map[string]interface{}{"code":20000, "data":users[token]}
	c.ServeJSON()
}

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Post() {
	c.Data["json"] = map[string]interface{}{"code":20000, "data":"success"}
	c.ServeJSON()
}

type ActivityController struct {
	beego.Controller
}

func (c *ActivityController) Get() {
	login := c.GetString("login")
	last_login := c.GetString("last_login")
	lastCount, count := statistic.GetStatisticActivityCount("lastLogin", "login", login, last_login)
	c.Data["json"] = map[string]interface{}{"code":20000, "data":map[string]interface{}{"last_active_count":lastCount,"active_count":count}}
	c.ServeJSON()
}

type NewlyController struct {
	beego.Controller
}

func (c *NewlyController) Get() {
	reg := c.GetString("reg")
	activy := c.GetString("active")
	reg_count, active_count := statistic.GetStatisticNewlyCount("regtime", "activeTime", reg, activy)
	//fmt.Println(reg_count, active_count)
	c.Data["json"] = map[string]interface{}{"code":20000, "data":map[string]interface{}{"reg_count":reg_count,"active_count":active_count}}
	c.ServeJSON()
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	day := c.GetString("day")
	register := statistic.GetUserStatistic(day, true)
	login := statistic.GetUserStatistic(day, false) + register
	total := statistic.GetUserCount()
	c.Data["json"] = map[string]interface{}{"code":20000, "data":map[string]interface{}{"active":login, "register":register, "total": total}}
	c.ServeJSON()
}

var tokens = map[string]map[string]string{
	"admin": {
		"token":"admin-token",
	},
	"editor": {
		"token":"editor-token",
	},
}

var users = map[string]map[string]interface{}{
	"admin-token": {
		"roles": []string{"admin"},
		"introduction": "I am a super administrator",
		"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name": "Super Admin",
  	},
	"editor-token": {
		"roles": []string{"editor"},
		"introduction": "I am an editor",
		"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name": "Normal Editor",
	},
}