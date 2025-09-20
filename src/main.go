package main

import (
	"go-challenge-timely-tag-system/app"
	"go-challenge-timely-tag-system/consts/envKeys"
	"go-challenge-timely-tag-system/jobs"
	"go-challenge-timely-tag-system/repository"
	"log"
	"os"
	"time"
)

func main() {
	time.Sleep(30 * time.Second)
	repository.Redis = repository.NewRedisRepository(os.Getenv(envKeys.RedisAddress), "", 0)
	go jobs.Init()
	app.StartApplication()
	log.Printf("app started")
}
