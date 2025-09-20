package jobs

import (
	"github.com/robfig/cron/v3"
	"go-challenge-timely-tag-system/consts/envKeys"
	"go-challenge-timely-tag-system/service"
	"log"
	"os"
	"time"
)

func Init() {
	c := cron.New()
	interval := os.Getenv(envKeys.RemoveOldRecordsInterval)
	log.Printf("interval:" + interval)
	_, err := c.AddFunc(interval, resetAllTags)
	if err != nil {

		log.Println(err.Error())

		return

	}
	c.Start()
	select {}
}

func resetAllTags() {
	log.Printf("removing old records at: %s\n", time.Now().UTC())
	_, err := service.RemoveOldSegmentUsers()
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
