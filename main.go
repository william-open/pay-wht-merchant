package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	adminRouters "likeadmin/admin/routers"
	admin "likeadmin/admin/service"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	genRouters "likeadmin/generator/routers"
	gen "likeadmin/generator/service"
	"likeadmin/middleware"
	"log"
	"net/http"
	"strconv"
	"time"
)

// initDI 初始化DI
// initDI 初始化DI
func initDI() {
	// 1. 获取admin模块的初始化函数集合
	regFunctions := admin.InitFunctions

	// 2. 追加gen模块的初始化函数
	regFunctions = append(regFunctions, gen.InitFunctions...)

	// 3. 注册数据库连接函数（使用标识符）
	// 注册主数据库
	if err := core.ProvideForDIWithName(core.DBMain, core.GetDB); err != nil {
		log.Fatalln("Failed to register main DB:", err)
	}

	// 注册订单数据库
	if err := core.ProvideForDIWithName(core.DBOrder, core.GetOrderDB); err != nil {
		log.Fatalln("Failed to register order DB:", err)
	}

	// 注册数据库获取函数（可以通过标识符获取任意数据库）
	if err := core.ProvideForDIWithName("databaseFactory", func(name string) (*gorm.DB, bool) {
		return core.GetDatabase(name)
	}); err != nil {
		log.Fatalln("Failed to register database factory:", err)
	}

	// 4. 遍历所有注册函数，进行依赖注入
	for i := 0; i < len(regFunctions); i++ {
		if err := core.ProvideForDI(regFunctions[i]); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("DI initialized with multiple database support")
}

// initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置静态路径
	router.Static(config.Config.PublicPrefix, config.Config.UploadDirectory)
	router.Static(config.Config.StaticPath, config.Config.StaticDirectory)
	// 设置中间件
	router.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	// 演示模式
	if config.Config.DisallowModify {
		router.Use(middleware.ShowMode())
	}
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 注册路由
	group := router.Group("/api")
	//core.RegisterGroup(group, routers.CommonGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.MonitorGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.SettingGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.SystemGroup, middleware.TokenAuth())

	routers := adminRouters.InitRouters[:]
	routers = append(routers, genRouters.InitRouters...)
	for i := 0; i < len(routers); i++ {
		core.RegisterGroup(group, routers[i])
	}
	return router
}

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 刷新日志缓冲
	defer func(Logger *zap.SugaredLogger) {
		err := Logger.Sync()
		if err != nil {

		}
	}(core.Logger)
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}
	// 订单库
	if core.GetOrderDB() != nil {
		db, _ := core.GetOrderDB().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}
	// 初始化DI
	initDI()
	// 初始化router
	router := initRouter()
	// 初始化server
	s := initServer(router)
	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())
}
