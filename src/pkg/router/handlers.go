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

func (r *Router) CreateSegmentHandler(c *gin.Context) {
	segment := &models.Segment{}

	// Перекладываем в структуру
	if err := c.ShouldBindJSON(&segment); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Проверяем название сегмента (начало названия: "AVITO_")
	SegmentName, err := parser.ParseSegmentName(segment.SegmentName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}

	// Добавляем новый сегмент
	segmentId, err := r.Db.CreateSegment(SegmentName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment created successfully", fmt.Sprintf("segment id: %d", segmentId): segment})
}

func (r *Router) DeleteSegmentHandler(c *gin.Context) {
	// Получаем название сегмента
	slug := c.Param("slug")

	deletedID, err := r.Db.DeleteSegment(slug)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment deleted successfully", "deleted_id": deletedID})
}

func (r *Router) DeleteUserHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя (число)
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	deletedID, err := r.Db.DeleteUser(userID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "deleted_id": deletedID})
}

func (r *Router) AddUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверка, что uid у пользователя (число)
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// проверка, что сегмент существует
	for _, segmentName := range segments.SegmentsName {
		if _, err = r.Db.GetIdSegment(segmentName); err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Segment: %s not found", segmentName)})
			return
		}
	}

	// добавление сегментов пользователю
	err = r.Db.CreateSegmentsUserRelation(userID, segments.SegmentsName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add segments to user"})
		return
	}

	response := models.UserSegmentsResponse{
		UserID:   userID,
		Segments: segments.SegmentsName,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segments added successfully", "response": response})
}

func (r *Router) DeleteUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверка, что uid у пользователя (число)
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// проверка, что сегмент существует
	for _, segmentName := range segments.SegmentsName {
		if _, err = r.Db.GetIdSegment(segmentName); err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Segment: %s not found", segmentName)})
			return
		}
	}

	// удаление сегментов у пользователя
	err = r.Db.DeleteSegmentsUserRelation(userID, segments.SegmentsName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segments to user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success to delete segments to user"})
}

func (r *Router) GetUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя (число)
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	id, err := r.Db.GetUserId(userID)
	if id == 0 || err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	segments, err := r.Db.GetUserSegments(id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segments to user"})
		return
	}
	response := models.UserSegmentsResponse{
		UserID:   userID,
		Segments: segments,
	}

	c.JSON(http.StatusOK, response)
}
