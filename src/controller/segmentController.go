package controller

import (
	"github.com/gin-gonic/gin"
	"go-challenge-timely-tag-system/models/request"
	"go-challenge-timely-tag-system/service"
)

func GetSegmentPopularity(context *gin.Context) {
	segment_id := context.Query("segment_id")

	result, err := service.GetSegmentUsersCount(segment_id)

	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(200, result)
}

func SubscribeSegment(context *gin.Context) {
	req := context.MustGet("req").(request.SubscribeUserRequest)
	err := service.SaveUserInSegment(req.UserId, req.SegmentId)
	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(200, gin.H{"message": "OK"})

}
