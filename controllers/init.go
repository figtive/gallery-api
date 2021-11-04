package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/controllers/v1"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/controllers/v1/middlewares"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(middlewares.GoogleOAuthMiddleware())
	{
		api.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, map[string]string{"ping": "pong"}) })
		apiV1 := api.Group("/v1")
		{
			auth := apiV1.Group("/auth")
			{
				auth.POST("/login", middlewares.AuthOnly(), v1.POSTLogin)
			}
			course := apiV1.Group("/course")
			{
				course.POST("/", v1.POSTCourse)
			}
			coursework := apiV1.Group("/coursework")
			{
				project := coursework.Group("/project")
				{
					project.POST("/", v1.POSTProject)
					project.GET("/", v1.GETProjects)
					project.GET("/:id", v1.GETProject)
					project.PUT("/thumbnail", v1.PUTThumbnail)
				}
				blog := coursework.Group("/blog")
				{
					blog.POST("/", v1.POSTBlog)
					blog.GET("/", v1.GETBlogs)
					blog.GET("/:id", v1.GETBlog)
				}
			}
			// team := apiV1.Group("/team")
			// {
			// 	team.POST("/", v1.POSTTeam)
			// }
		}

	}
	return router
}
