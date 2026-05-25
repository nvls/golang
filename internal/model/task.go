package model

import (
	"time"
)

type Task struct {
	Id         string
	Title      string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
