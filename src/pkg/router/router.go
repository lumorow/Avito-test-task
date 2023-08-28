package router

import (
	"Avito-test-task/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Rtr *gin.Engine
	Db  *postgres.Repo
}

func NewRouter(db *postgres.Repo) *Router {
	return &Router{
		Rtr: gin.Default(),
		Db:  db,
	}
}

func (r *Router) InitRoutes() {
	// Метод добавления сегмента
	r.Rtr.POST("/segment", r.CreateSegmentHandler)

	// Метод удаления сегмента
	r.Rtr.DELETE("/segment/:slug", r.DeleteSegmentHandler)

	// Метод добавления пользователя в сегменты
	r.Rtr.PUT("/user/:uid/segments", r.AddUserSegmentsHandler)

	// Метод удаления у пользователя сегменты
	r.Rtr.DELETE("/user/:uid/segments", r.DeleteUserSegmentsHandler)

	// Метод получения активных сегментов пользователя
	r.Rtr.GET("/user/:uid/segments", r.GetUserSegmentsHandler)

	// Метод удаления пользователя
	r.Rtr.DELETE("/user/:uid", r.DeleteUserHandler)
}
