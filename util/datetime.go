package util

import "time"

func DateToEpoch(inputTime time.Time) int64 {
	return inputTime.Unix()
}
