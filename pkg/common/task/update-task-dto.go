package task

import (
	"github.com/google/uuid"
	"time"
)

type UpdateTaskDto struct {
	ID        uuid.UUID         `json:"id"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	Url       string            `json:"url"`
	Headers   map[string]string `json:"headers"`
	Length    int               `json:"length"`
	Status    string            `json:"status"`
}
