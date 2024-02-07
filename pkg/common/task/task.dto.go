package task

import (
	"github.com/google/uuid"
	db "github.com/martirosharutyunyan/axxon-test-task/pkg/sqlc-gen"
	"time"
)

type TaskDto struct {
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Url       string        `json:"url"`
	Headers   []string      `json:"headers"`
	Length    int32         `json:"length"`
	Status    db.TaskStatus `json:"status"`
}

func NewTaskDto(taskEntity *db.Task) *TaskDto {
	return &TaskDto{
		ID:        taskEntity.ID,
		CreatedAt: taskEntity.CreatedAt,
		UpdatedAt: taskEntity.UpdatedAt,
		Url:       taskEntity.Url,
		Headers:   taskEntity.Headers,
		Length:    taskEntity.Length.Int32,
		Status:    taskEntity.Status,
	}
}
