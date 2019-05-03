package application

import (
	"github.com/gin-gonic/gin"
	"technopark_db/handlers"
)

type App struct {
	Router  *gin.Engine
	Handler *handlers.Handler
}

func (a *App) CreateRouter() (router *gin.Engine) {
	a.Router = gin.New()
	a.Router.Use(gin.Logger())
	a.Handler.Agregator.CreateTableAgr()
	// Создаем таблицы в бд
	api := a.Router.Group("/api")
	{
		api.GET("/CreateTables", a.Handler.CreateTables)
		// Чистим бд
		api.GET("/ClearTables", a.Handler.ClearTables)

		// Создание форума
		api.POST("/forum/:slug", a.Handler.CreateForum)
		// Создание ветки
		api.POST("/forum/:slug/create", a.Handler.CreateThread)

		api.GET("/forum/:slug/details", a.Handler.GetForum)

		api.GET("/forum/:slug/threads", a.Handler.GetThreads)

		//.Router.GET("/forum/:slug/users", a.Handler.CreatePost)

		//a.Router.GET("/post/:id/details", a.Handler.CreatePost)

		//a.Router.POST("/post/:id/details", a.Handler.CreatePost)

		api.POST("/service/clear", a.Handler.CreateTables)

		//a.Router.GET("/service/status", a.Handler.CreatePost)
		// Создание поста
		api.POST("/thread/:slug_or_id/create", a.Handler.CreatePost)
		// Получение информации о ветке
		//a.Router.GET("/thread/:slug_or_id/details", a.Handler.GetThread)
		//// Обновление информации о ветки
		//a.Router.POST("/thread/:slug_or_id/details", a.Handler.UpdateThreadDetails)
		// Получение постов из текущей ветки
		//a.Router.GET("/thread/:slug_or_id/posts", a.Handler.GetThreadPosts)
		// Голосование
		//a.Router.POST("/thread/:slug_or_id/vote", a.Handler.SetThreadVote)

		api.POST("/user/:nickname/create", a.Handler.CreateUser)

		api.GET("/user/:nickname/profile", a.Handler.GetUser)

		api.POST("/user/:nickname/profile", a.Handler.EditUser)
	}
	return
}
