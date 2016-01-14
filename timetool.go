package timetool

import (
	"log"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func GetStrTimeRangeNSeconds(startTime string, n_miniutes int64) (string, error) {
	t, err := time.Parse(TimeFormat, startTime)
	if err != nil {
		return "", err
	}
	newTime := t.Add(time.Minute * time.Duration(n_miniutes))
	result := newTime.Format(TimeFormat)
	if err != nil {
		log.Printf("format failed; err: %s startTime:%s; n_miniutes:%d \n", err, startTime, n_miniutes)
		return "", err
	}
	return result, nil
}

func GetStrTimeNowBeforNMinutes(n_miniutes int64) (string, error) {
	str := time.Now().Format(TimeFormat)
	return GetStrTimeRangeNSeconds(str, n_miniutes)
}
