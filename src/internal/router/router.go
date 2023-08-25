package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Rtr *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		Rtr: gin.Default(),
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

func (r *Router) CreateSegmentHandler(c *gin.Context) {

}

func (r *Router) DeleteSegmentHandler(c *gin.Context) {

}

func (r *Router) AddUserSegmentsHandler(c *gin.Context) {

}

func (r *Router) GetUserSegmentsHandler(c *gin.Context) {

}
