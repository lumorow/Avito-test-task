package router

import (
	"Avito-test-task/models"
	"Avito-test-task/parser"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// CreateSegmentHandler @Summary Create segment
// @Description Создание нового сегмента
// @Tags Segment
// @Accept json
// @Produce json
// @Param segment body models.SegmentRequest true "Данные сегмента"
// @Success 200 {object} models.SegmentResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /segment [post]
func (r *Router) CreateSegmentHandler(c *gin.Context) {
	segment := &models.SegmentRequest{}

	if err := c.BindJSON(&segment); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid JSON"})
		return
	}

	// Проверяем название сегмента (начало названия: "AVITO_")
	SegmentName, err := parser.ParseSegmentName(segment.SegmentName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: fmt.Sprintf("%s", err)})
		return
	}

	// Добавляем новый сегмент
	segmentId, err := r.Db.CreateSegment(SegmentName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to create segment"})
		return
	}

	response := models.SegmentResponse{
		Id:          segmentId,
		SegmentName: SegmentName,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSegmentHandler @Summary Delete segment
// @Description Удаление существующего сегмента
// @Tags Segment
// @Param slug path string true "Slug сегмента"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /segment/{slug} [delete]
func (r *Router) DeleteSegmentHandler(c *gin.Context) {
	slug := c.Param("slug")

	err := r.Db.DeleteSegment(slug)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to delete segment"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Segment deleted successfully"})
}

// DeleteUserHandler @Summary Delete user
// @Description Удаление пользователя
// @Tags User
// @Param uid path int true "ID пользователя"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{uid} [delete]
func (r *Router) DeleteUserHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя является числом
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	err = r.Db.DeleteUser(userID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "User deleted successfully"})
}

// AddUserSegmentsHandler @Summary Add user segments
// @Description Добавление сегментов пользователю
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "ID пользователя"
// @Param segments body models.Segments true "Данные сегментов"
// @Success 200 {object} models.UserSegmentsResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{uid}/segments [put]
func (r *Router) AddUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// проверка, что uid у пользователя является числом
	userUID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	// проверка, что сегмент существует и что его еще нет у пользователя
	for _, segmentName := range segments.SegmentsName {
		segmentID, err := r.Db.GetIdSegment(segmentName)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: fmt.Sprintf("Segment: %s not found", segmentName)})
			return
		}
		check, _ := r.Db.CheckSegmentUserRelation(userUID, segmentID)
		if check == true {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: fmt.Sprintf("segment: '%s' is already owned by the user with uid: %d", segmentName, userUID)})
			return
		}
	}

	// добавление сегментов пользователю
	err = r.Db.CreateSegmentsUserRelation(userUID, segments.SegmentsName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to insert segments to user"})
		return
	}

	response := models.UserSegmentsResponse{
		UserID:   userUID,
		Segments: segments.SegmentsName,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUserSegmentsHandler @Summary Delete user segments
// @Description Удаление сегментов у пользователя
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "ID пользователя"
// @Param segments body models.Segments true "Данные сегментов"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{uid}/segments [delete]
func (r *Router) DeleteUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// проверка, что uid у пользователя является числом
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	// проверка, что сегмент существует
	for _, segmentName := range segments.SegmentsName {
		if _, err = r.Db.GetIdSegment(segmentName); err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: fmt.Sprintf("Segment: %s not found", segmentName)})
			return
		}
	}

	// удаление сегментов у пользователя
	err = r.Db.DeleteSegmentsUserRelation(userID, segments.SegmentsName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to delete segments to user"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Success to delete segments to user"})
}

// GetUserSegmentsHandler @Summary Get user segments
// @Description Получение списка сегментов, в которые входит пользователь
// @Tags User
// @Param uid path int true "ID пользователя"
// @Success 200 {object} models.UserSegmentsResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{uid}/segments [get]
func (r *Router) GetUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя является числом
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	// получаем id пользователя
	id, err := r.Db.GetUserId(userID)
	if id == 0 || err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User not found"})
		return
	}

	segments, err := r.Db.GetUserSegments(id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to delete segments to user"})
		return
	}
	response := models.UserSegmentsResponse{
		UserID:   userID,
		Segments: segments,
	}

	c.JSON(http.StatusOK, response)
}
