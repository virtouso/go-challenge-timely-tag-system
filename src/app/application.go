package app

import (
	"github.com/gin-gonic/gin"
	"go-challenge-timely-tag-system/consts/envKeys"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	port := ":" + os.Getenv(envKeys.Port)
	router.Run(port)

}
