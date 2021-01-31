package entities

import "time"

type Appointment struct {
	Entity
	UserId          string    `db:"UserId"`
	MessageId       int       `db:"MessageId"`
	SchedulingStart time.Time `db:"SchedulingStart"`
	SchedulingEnd   time.Time `db:"SchedulingEnd"`
	ExamName        string    `db:"ExamName"`
	ExamValue       string    `db:"ExamValue"`
}
