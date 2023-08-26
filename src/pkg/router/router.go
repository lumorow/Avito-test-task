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
	r.Rtr.POST("/segment", r.CreateSegmentHandler)

	// Метод удаления сегмента
	r.Rtr.DELETE("/segment/:slug", r.DeleteSegmentHandler)

	// Метод добавления пользователя в сегмент
	r.Rtr.PUT("/user/:id/segments", r.AddUserSegmentsHandler)

	// Метод получения активных сегментов пользователя
	r.Rtr.GET("/user/:id/segments", r.GetUserSegmentsHandler)
}
