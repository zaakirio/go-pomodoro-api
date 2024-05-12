package models

import "time"

type Pomodoro struct {
	ID        string    `firestore:"-" json:"id"`
	StartTime time.Time `json:"startTime"`
	Duration  int       `json:"duration"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
}

func NewPomodoro(task string, duration int) *Pomodoro {
	return &Pomodoro{
		StartTime: time.Now(),
		Duration:  duration,
		Task:      task,
		Completed: false,
	}
}
