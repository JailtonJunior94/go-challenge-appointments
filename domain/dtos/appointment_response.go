package dtos

import "time"

type AppointmentResponse struct {
	Id              string    `json:"id"`
	ExamName        string    `json:"examName"`
	ExamValue       float64   `json:"examValue"`
	SchedulingStart time.Time `json:"schedulingStart"`
	SchedulingEnd   time.Time `json:"schedulingEnd"`
	Active          bool      `json:"active"`
}
