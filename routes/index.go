package routes

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	mdl "hankliu.com.cn/go-my-blog/middleware"
	"hankliu.com.cn/go-my-blog/services"
)

var (
	// SBlogService blog service struct instance
	SBlogService = &services.BlogService{}
)

// InitialRoutes inittail routers
func InitialRoutes(router *gin.Engine) {
	initialFrontendRoutes(router)
	initialBackendRoutes(router)
}

func actionHandler(hander gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		hander(c)
	}
}

func initialBackendRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/login", actionHandler(SBlogService.Login))
		api.POST("/register", actionHandler(SBlogService.Register))
		api.GET("/blog/:id", actionHandler(SBlogService.GetPost))
		api.GET("/blogs", actionHandler(SBlogService.SearchPosts))
		api.POST("/blogs/:id/comment", actionHandler(SBlogService.AppendComment))

		api.PUT("/updatepassword", mdl.AuthLogin(actionHandler(SBlogService.UpdatePassword)))
		api.POST("/blog", mdl.AuthLogin(actionHandler(SBlogService.CreatePost)))
		api.PUT("/blog/:id", mdl.AuthLogin(actionHandler(SBlogService.UpdatePost)))
		api.DELETE("/blog/:id", mdl.AuthLogin(actionHandler(SBlogService.DeletePost)))
		api.GET("/own/blogs", mdl.AuthLogin(actionHandler(SBlogService.SearchOwnPosts)))
		api.DELETE("/blogs/:id/comment/:commentId", mdl.AuthLogin(actionHandler(SBlogService.RemoveComment)))
	}
}

func initialFrontendRoutes(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	router.GET("/index", func(c *gin.Context) {
		render(c, "index", nil)
	})
}

func getStaticFiles(name string) (string, gin.H) {
	data := gin.H{}
	data["csses"] = []interface{}{map[string]string{"src": "/static/stylesheets/" + name + "/index.css"}}
	data["jses"] = []interface{}{map[string]string{"src": "/static/javascripts/" + name + "/index.js"}}
	return name + ".html", data
}

func render(c *gin.Context, name string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}

	viewName, commonData := getStaticFiles(name)
	for key, value := range commonData {
		data[key] = value
	}

	c.HTML(http.StatusOK, viewName, data)
}
