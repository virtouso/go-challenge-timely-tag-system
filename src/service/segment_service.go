package service

import (
	"go-challenge-timely-tag-system/consts/envKeys"
	"go-challenge-timely-tag-system/repository"
	"os"
	"strconv"
	"time"
)

func SaveUserInSegment(userId string, segmentId string) (err error) {

	time := time.Now().UTC()
	unixTimeSeconds := time.Unix()
	result := repository.Redis.AddToZSet(segmentId, userId, float64(unixTimeSeconds))

	return result.Err()
}

func GetSegmentUsersCount(segmentId string) (count int64, err error) {
	result, err := repository.Redis.CountZSet(segmentId)
	return result, err
}

func RemoveOldSegmentUsers() (removedCount int64, err error) {
	minutesString := os.Getenv(envKeys.KeepAliveMinutes)
	minutesCount, err := strconv.Atoi(minutesString)

	time := time.Now().UTC().Add(-time.Duration(minutesCount) * time.Minute)
	unixTimeSeconds := time.Unix()
	count, err := repository.Redis.RemoveBelowScoreAll(float64(unixTimeSeconds))

	return count, err
}
