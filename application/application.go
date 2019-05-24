package application

import (
	"github.com/Pickausernaame/technopark_db/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

type App struct {
	Router  *gin.Engine
	Handler *handlers.Handler
}

func CreateApp(conn *pgx.ConnPoolConfig) *App {
	var a App
	a.Handler = handlers.CreateHandler(conn)
	return &a
}

func (a *App) CreateRouter() (router *gin.Engine) {
	a.Router = gin.New()
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())
	//a.Handler.Agregator.ClearTableAgr()

	a.Handler.Agregator.CreateTableAgr()
	a.Handler.Agregator.CreateIndexes()

	// Создаем таблицы в бд
	api := a.Router.Group("/api")
	{
		api.GET("", a.Handler.Connected)
		api.GET("/CreateTables", a.Handler.CreateTables)
		// Чистим бд
		api.GET("/ClearTables", a.Handler.ClearTables)

		// Создание форума
		api.POST("/forum/:slug", a.Handler.CreateForum)
		// Создание ветки
		api.POST("/forum/:slug/create", a.Handler.CreateThread)

		api.GET("/forum/:slug/details", a.Handler.GetForum)

		api.GET("/forum/:slug/threads", a.Handler.GetThreads)

		api.GET("/forum/:slug/users", a.Handler.GetForumUsers)

		api.GET("/post/:id/details", a.Handler.GetPostDetails)

		api.POST("/post/:id/details", a.Handler.UpdatePost)

		api.POST("/service/clear", a.Handler.ClearTables)

		api.GET("/service/status", a.Handler.ServiceStatus)
		// Создание поста
		api.POST("/thread/:slug_or_id/create", a.Handler.CreatePost)
		// Получение информации о ветке
		api.GET("/thread/:slug_or_id/details", a.Handler.GetThreadDetails)
		//// Обновление информации о ветки
		api.POST("/thread/:slug_or_id/details", a.Handler.UpdateThreadDetails)
		// Получение постов из текущей ветки
		api.GET("/thread/:slug_or_id/posts", a.Handler.GetThreadPosts)
		// Голосование
		api.POST("/thread/:slug_or_id/vote", a.Handler.SetThreadVote)

		api.POST("/user/:nickname/create", a.Handler.CreateUser)

		api.GET("/user/:nickname/profile", a.Handler.GetUser)

		api.POST("/user/:nickname/profile", a.Handler.EditUser)
	}
	return
}
