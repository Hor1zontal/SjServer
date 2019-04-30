package http

import (
	"aliens/log"
	"context"
	"gangbu/exception"
	"gangbu/module/game/config"
	"gangbu/module/game/http/api"
	"gangbu/module/game/http/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

var HttpSrv *http.Server

func Init() {
	if config.Server.HTTPAddress == "" {
		return
	}
	if !log.DEBUG {
		//gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.Use(Recovery(), gin.Logger())

	if config.Server.IsCors {
		r.Use(Cors())
	}

	//log.Debug("mode:%v",gin.Mode())

	r.GET("/test", Test)
	r.GET("/delete", api.DeleteUser)

	r.GET("/notice/public", api.PubNotice)						//发布公告

	r.GET("/statistic/newly", api.NewlyRemain) //新增留存
	r.GET("/statistic/active", api.ActivityRemain) //活跃留存

	auth := r.Group("")
	auth.Use(api.SignCheck)

	auth.GET("/users", api.Login)
	auth.Use(api.JWTAuth)
	{
		//---------------------GET(SELECT)---------------------//
		auth.GET("/roles", api.GetRoleInfo) 						//获取角色数据(登录成功后第一个调的接口)
		auth.GET("/roles/ranks", api.GetRankings) 				//获取排行榜
		auth.GET("/roles/rank", api.GetRoleRank) 					//获取用户排名
		auth.GET("/roles/rank/limit", api.GetScoreRank)

		auth.GET("/game",api.GetGame) 							//获取游戏中数据
		auth.GET("/notice",api.GetNotice)
		auth.GET("/items", api.GetItems)          				//获取物品

		/*---------------------POST(CREATE)---------------------*/
		auth.POST("/game", api.EnterGame) 						//进入游戏 (使用体力，携带道具)

		/*---------------------PUT(UPDATE)----------------------*/
		auth.PUT("/users",api.UpdateUser)                  		//更新用户名和头像地址
		auth.PUT("/roles/energy", api.UpdateEnergy) 				//更新体力
		auth.PUT("/roles/guide", api.PassGuide) 					//通过新手引导 更新

		auth.PUT("/game", api.UpdateGame)
		//auth.PUT("/box", api.UpdateBoxData)						//更新宝箱数据

		auth.PUT("/props/help", api.AddHelpProp)					// 帮助
		auth.PUT("/items/add",api.AddItem)						// 新增物品
		auth.PUT("/items/use", api.UseItem)						// 使用物品

		/*---------------------PATCH(UPDATE)---------------------*/
		/*---------------------DELETE(DELETE)--------------------*/
	}

	HttpSrv = &http.Server{
		Addr: config.Server.HTTPAddress,
		Handler:r,
	}

	go func() {
		log.Debug("Http Bind Port %v", HttpSrv.Addr)
		if err := HttpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("start http service err : %v", err)
		}
	}()

}

func Close() {
	log.Debug("Shutdown Server ...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := HttpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:",err)
	}
	log.Debug("Server exiting")
}

func Test(c *gin.Context) {
	//time.Sleep(3 * time.Second)
	//cCp := c.Copy()
	//go func() {
	helper.ResponseWithCode(c, exception.CodeSuccess)
	//}()

}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//log.Debug("error:%v",err)
				switch err.(type) {
				case exception.ErrorCode:
					helper.ResponseWithCode(c, err.(exception.ErrorCode))
				default:
					log.Error("%v - %v", err, string(debug.Stack()))
					log.Error("%v", err)
					debug.PrintStack()
					helper.ResponseWithCode(c, exception.InternalError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token, sign")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}