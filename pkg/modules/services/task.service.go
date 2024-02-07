package services

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/common/task"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/database"
	dbErrors "github.com/martirosharutyunyan/axxon-test-task/pkg/database/errors"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/sqlc-gen"
	"time"
)

type taskService struct{}

type ITaskService interface {
	Create(createTaskDto *task.CreateDto) (taskEntity *db.Task, err error)
	GetById(id uuid.UUID) (taskEntity *db.Task, err error)
}

var _ ITaskService = &taskService{}

func (t taskService) Create(createTaskDto *task.CreateDto) (taskEntity *db.Task, err error) {
	ctx := context.Background()
	instance := db.New(database.Instance)
	taskEntity, err = instance.Create(ctx, createTaskDto.URL)

	if err != nil {
		return nil, dbErrors.Parse(err)
	}

	go func(createTaskDto *task.CreateDto) {
		client := req.C()
		client.SetTimeout(time.Second * 5)
		request := client.R()
		request.Method = createTaskDto.Method

		for key, value := range createTaskDto.Headers {
			request.Headers.Set(key, value)
		}
		res, err := request.Get(createTaskDto.URL)

		if err != nil {
			instance.Update(ctx, db.UpdateParams{
				ID:     taskEntity.ID,
				Status: db.TaskStatusERROR,
			})
			return
		}

		var headers []string
		for key, _ := range res.Header {
			headers = append(headers, key)
		}

		if err != nil {
			instance.Update(ctx, db.UpdateParams{
				ID:     taskEntity.ID,
				Status: db.TaskStatusERROR,
			})
			return
		}

		instance.Update(ctx, db.UpdateParams{
			ID:      taskEntity.ID,
			Headers: headers,
			Status:  db.TaskStatusDONE,
			Length:  sql.NullInt32{Int32: int32(len(res.Bytes())), Valid: true},
		})
	}(createTaskDto)

	return taskEntity, nil
}

func (t taskService) GetById(id uuid.UUID) (taskEntity *db.Task, err error) {
	instance := db.New(database.Instance)
	ctx := context.Background()

	taskEntity, err = instance.GetById(ctx, id)

	return taskEntity, dbErrors.Parse(err)
}

func NewTaskService() ITaskService {
	return taskService{}
}
