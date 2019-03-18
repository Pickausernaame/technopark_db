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
	// Создаем таблицы в бд
	a.Router.GET("/CreateTables", a.Handler.CreateTables)
	// Чистим бд
	a.Router.GET("/ClearTables", a.Handler.ClearTables)

	// Создание форума
	a.Router.POST("/forum/:slug", a.Handler.CreateForum)
	// Создание ветки
	a.Router.POST("/forum/:slug/create", a.Handler.CreateThread)
	// Создание поста/
	a.Router.POST("/thread/:slug_or_id/create", a.Handler.CreatePost)
	// Получение информации о ветке
	a.Router.GET("/thread/:slug_or_id/details", a.Handler.GetThreadDetails)
	// Обновление информации о ветки
	a.Router.POST("/thread/:slug_or_id/details", a.Handler.UpdateThreadDetails)
	// Получение постов из текущей ветки
	a.Router.GET("/thread/:slug_or_id/posts", a.Handler.GetThreadPosts)
	// Голосование
	a.Router.POST("/thread/:slug_or_id/vote", a.Handler.SetThreadVote)

	return
}
