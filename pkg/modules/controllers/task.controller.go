package controllers

import (
	"github.com/google/uuid"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/common/task"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/modules/response"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/modules/services"
	"net/http"
)

type taskController struct {
	taskService services.ITaskService
}

func (c taskController) Create(ctx *response.Context) {
	createTaskDto := new(task.CreateDto)

	if err := ctx.ShouldBindJSON(createTaskDto); err != nil {
		ctx.UnprocessableEntity(err.Error())
		return
	}

	taskEntity, err := c.taskService.Create(createTaskDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]uuid.UUID{"id": taskEntity.ID})
}

func (c taskController) GetById(ctx *response.Context) {
	taskIdString := ctx.Param("id")
	taskId, err := uuid.Parse(taskIdString)

	if err != nil {
		ctx.UnprocessableEntity(err.Error())
		return
	}

	taskEntity, err := c.taskService.GetById(taskId)

	if err != nil {
		ctx.Error(err)
		return
	}

	taskDto := task.NewTaskDto(taskEntity)
	ctx.JSON(http.StatusOK, taskDto)
}

func newTaskController(taskService services.ITaskService) *taskController {
	return &taskController{taskService: taskService}
}

func InitTaskController() {
	taskGroup := ApiRouter.Group("/task")

	taskService := services.NewTaskService()
	taskController := newTaskController(taskService)

	taskGroup.POST("/", response.Handler(taskController.Create))
	taskGroup.GET("/:id", response.Handler(taskController.GetById))
}
