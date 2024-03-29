package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/controllers/v1"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/controllers/v1/middlewares"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
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
				bookmark.POST("/:coursework_id", middlewares.AuthOnly(), v1.POSTBookmark)
				bookmark.GET("/:coursework_id", middlewares.AuthOnly(), v1.GETBookmarkStatus)
			}
			course := apiV1.Group("/course")
			{
				course.POST("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.POSTCourse)
				course.GET("", v1.GETCourses)
				course.PUT("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.PUTCourse)
				course.GET("/:course_id", v1.GETCourse)
				course.DELETE("/:course_id", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.DELETECourse)
			}
			coursework := apiV1.Group("/coursework")
			{
				project := coursework.Group("/project")
				{
					project.PUT("/thumbnail", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.PUTThumbnail)
					project.DELETE("/thumbnail", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.DELETEThumbnail)
					project.POST("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.POSTProject)
					project.PUT("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.PUTProject)
					project.GET("", v1.GETProjects)
					project.GET("/:course_id", v1.GETProjects)
					project.GET("/:course_id/:coursework_id", v1.GETProject)
					project.DELETE("/:course_id/:project_id", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.DELETEProject)
				}
				blog := coursework.Group("/blog")
				{
					blog.POST("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.POSTBlog)
					blog.PUT("", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.PUTBlog)
					blog.GET("", v1.GETBlogs)
					blog.GET("/:course_id", v1.GETBlogs)
					blog.GET("/:course_id/:coursework_id", v1.GETBlog)
					blog.DELETE("/:course_id/:blog_id", middlewares.AuthOnly(), middlewares.AdminOnly(), v1.DELETEBlog)
				}
			}
			vote := apiV1.Group("/vote")
			{
				vote.GET("/quota", middlewares.AuthOnly(), v1.GETVoteQuota)
				vote.GET("/project", middlewares.AuthOnly(), v1.GETVotedProject)
				vote.GET("/blog", middlewares.AuthOnly(), v1.GETVotedBlogs)
				vote.GET("/count/:coursework_id", v1.GETVoteCount)
				vote.POST("/:coursework_id", middlewares.AuthOnly(), v1.POSTVote)
				vote.GET("/:coursework_id", middlewares.AuthOnly(), v1.GETVoteStatus)
			}
			leaderboard := apiV1.Group("/leaderboard")
			{
				leaderboard.GET("/:course_id/blog", v1.GETBlogLeaderboard)
				leaderboard.GET("/:course_id/project", v1.GETProjectLeaderboard)
			}
		}

	}
	return router
}
