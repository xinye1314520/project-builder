package route

import (
	"github.com/gin-gonic/gin"
	"project-builder/util"
	"net/http"
	"project-builder/controller"
	"path/filepath"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	//sessionName := utils.Config.Section("session").Key("session.name").String()
	//domain := utils.Config.Section("session").Key("session.domain").String()
	//maxAge, _ := utils.Config.Section("session").Key("session.gc.max.life.time").Int()
	//csrfSecret := utils.Config.Section("session").Key("csrf.secret").String()
	////Session配置，Redis存储
	////store,err:=redis.NewStore(10,"tcp","rs1.rs.baidu.com:6379","",[]byte("asfajfa;lskjr2"))
	//store, err := redis.NewStoreWithPool(utils.RedisPool, []byte("as&8(0fajfa;lskjr2"))
	//store.Options(sessions.Options{
	//	"/",
	//	domain,
	//	maxAge,
	//	false, //https 时使用
	//	true,  //true:JS脚本无法获取cookie信息
	//})
	//
	//if err != nil {
	//	// Handle the error. Probably bail out if we can't connect.
	//	fmt.Println("redis.NewStore error")
	//}
	//
	//Router.Use(sessions.Sessions(sessionName, store))

	//Router.Use(csrf.Middleware(csrf.Options{
	//	Secret: csrfSecret,
	//	ErrorFunc: func(c *gin.Context) {
	//		c.String(400, "CSRF token mismatch")
	//		c.Abort()
	//	},
	//}))

}

func SetupRouter() *gin.Engine {

	//静态目录配置
	public := util.Config.Section("router").Key("public").String()
	Router.Static("/public", public)

	//模板
	viewPath := util.Config.Section("router").Key("view.path").String()
	Router.LoadHTMLGlob(filepath.Join(util.GetProPath(), viewPath))

	//用户route
	Router.GET("/user/list", controller.QueryUserList)
	Router.POST("/user/add", controller.UserAdd)
	Router.GET("/user/add", controller.UserAdd)
	//Router.GET("/user/show", users.UserShowAction)
	//Router.GET("/user/edit/:id", users.UserEditAction)

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})
	return Router
}
