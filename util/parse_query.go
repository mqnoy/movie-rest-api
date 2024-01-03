package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func ParseQueryToInt(query string) int {
	result, err := strconv.Atoi(query)
	if err != nil {
		return -1
	}

	return result
}

func ParseQueryToFloat(query string) float64 {
	result, err := strconv.ParseFloat(query, 64)
	if err != nil {
		return -1
	}

	return result
}

func ParseQueryToString(query string) string {
	return query
}

func ParseQueryToBool(query string) *bool {
	result, err := strconv.ParseBool(query)
	if err != nil {
		return nil
	}

	return &result
}

func ParseQueryToTime(query string) time.Time {
	result, err := time.Parse("2006-01-02 15:04:05", query)
	if err != nil {
		return time.Time{}
	}

	return result
}

func ParseQueryToDate(query string) time.Time {
	result, err := time.Parse("2006-01-02", query)
	if err != nil {
		return time.Time{}
	}

	return result
}

func ParseQueryToStringArray(query string) []string {
	var strings []string

	err := json.Unmarshal([]byte(query), &strings)
	if err != nil {
		fmt.Println("Error parsing query:", err)
		return []string{}
	}

	return strings
}
