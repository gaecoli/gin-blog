package router

import (
	"gin-blog/internal/handle"
	"gin-blog/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// handle instance
var (
	blogViewApi handle.BlogInfo // blog 前台展示 handle 处理函数
	articleApi  handle.Article  // 文章相关 handle 处理函数
	categoryApi handle.Category // 分类相关 handle 处理函数
	tagApi      handle.Tag
	authApi     handle.LoginApi
	userApi     handle.User
)

func RegisterRouter(r *gin.Engine) {
	registerJwtHandler(r)
	registerBlogViewHandler(r)
	registerBlogManagerHandler(r)
}

// blog 的登录，注册接口，不需要鉴权
func registerJwtHandler(r *gin.Engine) {
	blog := r.Group("/api/v1")

	blog.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})

	blog.POST("/login", authApi.Login)
	blog.GET("/logout", authApi.Logout)
	blog.POST("/register", authApi.Register)
	blog.GET("/send_code", authApi.SendEmailCode)

}

// blog 管理接口，全部需要登录和鉴权
func registerBlogManagerHandler(r *gin.Engine) {
	auth := r.Group("/api/v1/manager")

	auth.Use(middleware.JwtAuthMiddleware())
	// TODO: 获取管理员编辑的操作日志

	// TODO: 补齐 about 编辑和获取
	about := auth.Group("/setting")
	{
		about.GET("/about", func(c *gin.Context) {
			c.JSON(200, "获取 about 页面信息")
		})
		about.PUT("/about", func(c *gin.Context) {
			c.JSON(200, "更新 about 页面")
		})
	}

	user := auth.Group("/user")
	{
		user.GET("/list", userApi.GetUserInfoList)
		user.GET("/:id", userApi.GetUserInfoById)
		user.PUT("/archive", userApi.UpdateUserDisableInfo)
		user.POST("", userApi.UpdateUserInfo)
		user.POST("/change_password", userApi.UpdateUserPassword)
	}

	category := auth.Group("/category")
	{
		category.GET("/list", categoryApi.GetCategoryList)
		category.POST("/new", categoryApi.CreateCategory)
		category.POST("/update", categoryApi.UpdateCategory)
		category.DELETE("", categoryApi.DeleteCategory)
	}

	tag := auth.Group("/tag")
	{
		tag.GET("/list", tagApi.GetTagList)
		tag.POST("", tagApi.CreateOrUpdateTag)
		tag.DELETE("", tagApi.DeleteTag)
	}

	article := auth.Group("/article")
	{
		article.GET("/list", articleApi.GetArticleList)
		article.POST("", articleApi.CreateArticle)
		article.GET("/:id", articleApi.GetArticle)
		article.PUT("archive", articleApi.SoftDeleteArticle)
		article.DELETE("", articleApi.DeleteArticle)
	}

}

// blog view handler, don't need to auth
func registerBlogViewHandler(r *gin.Engine) {
	blog := r.Group("/api/v1")

	//blog.GET("/about", )
	blog.GET("/home", blogViewApi.GetHomeInfo)

	article := blog.Group("/article")
	{
		article.GET("/list", articleApi.GetArticleList)
		article.GET("/:id", articleApi.GetArticle)
	}

	category := blog.Group("/category")
	{
		category.GET("/list", categoryApi.GetCategoryList)
	}

	tag := blog.Group("/tag")
	{
		tag.GET("/list", tagApi.GetTagList)
	}
}
