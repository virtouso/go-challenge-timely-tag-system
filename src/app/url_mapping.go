package app

import (
	"go-challenge-timely-tag-system/controller"
	"go-challenge-timely-tag-system/middleware"
	"go-challenge-timely-tag-system/models/request"
)

func mapUrls() {
	router.GET("api/v1/get_segment_popularity", controller.GetSegmentPopularity)
	router.POST("api/v1/subscribe_segment", middleware.BindJSON[request.SubscribeUserRequest](), controller.SubscribeSegment)
}
