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

// Создание сегмента

func (r *Router) CreateSegmentHandler(c *gin.Context) {
	segment := &models.Segment{}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create segment"})
		return
	}

	response := models.Segment{
		Id:          segmentId,
		SegmentName: SegmentName,
	}

	c.JSON(http.StatusOK, response)
}

// Удаление сегмента

func (r *Router) DeleteSegmentHandler(c *gin.Context) {
	slug := c.Param("slug")

	err := r.Db.DeleteSegment(slug)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment deleted successfully"})
}

// Удаление пользователя

func (r *Router) DeleteUserHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя является числом
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = r.Db.DeleteUser(userID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Добавляем сегменты пользователю

func (r *Router) AddUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверка, что uid у пользователя является числом
	userUID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// проверка, что сегмент существует и что его еще нет у пользователя
	for _, segmentName := range segments.SegmentsName {
		segmentID, err := r.Db.GetIdSegment(segmentName)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Segment: %s not found", segmentName)})
			return
		}
		check, _ := r.Db.CheckSegmentUserRelation(userUID, segmentID)
		if check == true {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("segment: '%s' is already owned by the user with uid: %d", segmentName, userUID)})
			return
		}
	}

	// добавление сегментов пользователю
	err = r.Db.CreateSegmentsUserRelation(userUID, segments.SegmentsName)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
		return
	}

	response := models.UserSegmentsResponse{
		UserID:   userUID,
		Segments: segments.SegmentsName,
	}

	c.JSON(http.StatusOK, response)
}

// Удаление сегментов у пользователя

func (r *Router) DeleteUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")
	segments := &models.Segments{}
	if err := c.ShouldBindJSON(segments); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверка, что uid у пользователя является числом
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete segments to user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success to delete segments to user"})
}

// Получение сегментов, в которые входит пользователь

func (r *Router) GetUserSegmentsHandler(c *gin.Context) {
	UID := c.Param("uid")

	// проверка, что uid у пользователя является числом
	userID, err := strconv.Atoi(UID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// получаем id пользователя
	id, err := r.Db.GetUserId(userID)
	if id == 0 || err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	segments, err := r.Db.GetUserSegments(id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete segments to user"})
		return
	}
	response := models.UserSegmentsResponse{
		UserID:   userID,
		Segments: segments,
	}

	c.JSON(http.StatusOK, response)
}
