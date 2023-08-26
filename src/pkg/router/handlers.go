package router

import (
	"Avito-test-task/models"
	"Avito-test-task/parser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) CreateSegmentHandler(c *gin.Context) {
	segment := &models.Segment{}

	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	SegmentName, err := parser.ParseSegmentName(segment.SegmentName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	segmentId, err := r.Db.CreateSegment(SegmentName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment created successfully", fmt.Sprintf("segment id: %d", segmentId): segment})
}

func (r *Router) DeleteSegmentHandler(c *gin.Context) {

}

func (r *Router) AddUserSegmentsHandler(c *gin.Context) {

}

func (r *Router) GetUserSegmentsHandler(c *gin.Context) {

}
