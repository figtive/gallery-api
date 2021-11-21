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
				auth.POST("/login", middlewares.LooseAuthOnly(), v1.POSTLogin)
			}
			bookmark := apiV1.Group("/bookmark")
			{
				bookmark.GET("/blog", middlewares.AuthOnly(), v1.GETBookmarkBlogs)
				bookmark.GET("/project", middlewares.AuthOnly(), v1.GETBookmarkProjects)
				bookmark.POST("/:courseworkId", middlewares.AuthOnly(), v1.POSTBookmark)
				bookmark.GET("/:courseworkId", middlewares.AuthOnly(), v1.GETBookmarkStatus)
			}
			course := apiV1.Group("/course")
			{
				course.POST("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.POSTCourse)
			}
			coursework := apiV1.Group("/coursework")
			{
				project := coursework.Group("/project")
				{
					project.POST("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.POSTProject)
					project.GET("", v1.GETProjects)
					project.PUT("/thumbnail", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.PUTThumbnail)
					project.GET("/:id", v1.GETProject)
				}
				blog := coursework.Group("/blog")
				{
					blog.POST("", middlewares.AdminOnly(), v1.POSTBlog)
					blog.GET("", v1.GETBlogs)
					blog.GET("/:id", v1.GETBlog)
				}
			}
			vote := apiV1.Group("/vote")
			{
				vote.GET("/quota", middlewares.AuthOnly(), v1.GETVoteQuota)
				vote.GET("/count/:courseworkId", v1.GETVoteCount)
				vote.POST("/:courseworkId", middlewares.AuthOnly(), v1.POSTVote)
				vote.GET("/:courseworkId", middlewares.AuthOnly(), v1.GETVoteStatus)
			}
			leaderboard := apiV1.Group("/leaderboard")
			{
				leaderboard.GET("/:courseId/blog", v1.GETBlogLeaderboard)
				leaderboard.GET("/:courseId/project", v1.GETProjectLeaderboard)
			}
		}

	}
	return router
}
