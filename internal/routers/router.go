package routers

import (
	"net/http"
	"time"

	"github.com/MyBlog/global"
	"github.com/MyBlog/internal/middleware"
	"github.com/MyBlog/pkg/limiter"

	_ "github.com/MyBlog/docs"
	"github.com/MyBlog/internal/routers/api"
	v1 "github.com/MyBlog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	// 初始化Engine实例
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		// 引入中间件
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	// 引入中间件
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	// r.GET将定义的路由注册进去
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", api.GetAuth)

	// 返回路由组
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)       //增加标签
		apiv1.DELETE("/tags/:id", tag.Delete) //删除标签
		apiv1.PUT("/tags/:id", tag.Update)    //更改标签
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List) //查询标签

		apiv1.POST("/articles", article.Create)       //添加文章
		apiv1.DELETE("/articles/:id", article.Delete) //删除文章
		apiv1.PUT("/articles/:id", article.Update)    //更改文章
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get) //查询文章
		apiv1.GET("/articles", article.List)    //查询文章列表
	}

	return r
}
