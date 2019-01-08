package route

import (
	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"golang_iris_xorm/application/common"
	"golang_iris_xorm/config"
	"golang_iris_xorm/extend"
)

// 路由初始化
func RouteInit(app *iris.Application) {

	// 404处理
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(common.JsonReturn(404, "找不到页面", nil))
		return
	})

	// CRS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	// 控制器实例化


	// 子域名绑定
	api := app.Party("api.", crs, LimitHandler()).AllowMethods(iris.MethodOptions)
	{


		// ------------------------------------------
		// -----------此分割线以下接口需要登录------------
		// ------------------------------------------
		api.Use(new(extend.Jwt).CheckLogin)


	}
}

// http限流
func LimitHandler() context.Handler {
	// 限制HTTP请求(每秒5次)
	limiter := tollbooth.NewLimiter(config.Limiter, nil)

	return func(ctx context.Context) {
		httpError := tollbooth.LimitByRequest(limiter, ctx.ResponseWriter(), ctx.Request())
		if httpError != nil {
			// json报错
			ctx.JSON(common.JsonReturn(httpError.StatusCode, httpError.Message, nil))
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}