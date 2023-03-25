package structs

import "time"

type MetaData struct {
	Id              int       `json:"id"`
	StartTime       time.Time `json:"start_date"`
	Sport           string    `json:"type"`
	AveragePower    float64   `json:"weighted_average_watts"`
	AverageSpeed    float64   `json:"average_speed"`
	AverageHearRate float64   `json:"average_heartrate"`
	TotalTime       int       `json:"elapsed_time"`
	TotalDistance   float64   `json:"distance"`
	Page            int       `json:"page"`
}
