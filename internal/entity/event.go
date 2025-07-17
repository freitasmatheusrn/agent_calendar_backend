package entity

import "time"

type Event struct {
	Summary     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

func NewEvent(summary, description string, starTime, endTime time.Time)(*Event, error){
	return &Event{
		Summary: summary,
		Description: description,
		StartTime: starTime,
		EndTime: endTime,
	},nil
}