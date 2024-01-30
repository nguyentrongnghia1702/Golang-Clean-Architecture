package usecase

import (
	"context"
	"errors"
	"mcs-nghiadeptrai/mcs-common/logger"
	"mcs-nghiadeptrai/mcs-todo-item/domain"
	"time"
)

type todoUsecase struct {
	taskRepo       domain.ToDoItemRepository
	contextTimeout time.Duration
}

func NewTodoUsecase(taskRepo domain.ToDoItemRepository, timeout time.Duration) domain.TodoItemUseCase {
	return &todoUsecase{
		taskRepo:       taskRepo,
		contextTimeout: timeout,
	}
}

// CreateItem implements domain.TodoItemUseCase.
func (useCase *todoUsecase) CreateItem(ctx context.Context, data *domain.ToDoItem) error {
	if data.Title == "" {
		logger.LogErrorNoContext("title can not be blank")
		return errors.New("title can not be blank")
	}
	return useCase.taskRepo.CreateItem(ctx, data)
}

// DeleteItem implements domain.TodoItemUseCase.
func (useCase *todoUsecase) DeleteItem(ctx context.Context, conditions map[string]interface{}) error {
	return useCase.taskRepo.DeleteItem(ctx, conditions)
}

// GetAllItem implements domain.TodoItemUseCase.
func (useCase *todoUsecase) GetAllItem(ctx context.Context, paging *domain.DataPaging) ([]domain.ToDoItem, error) {
	return useCase.taskRepo.GetAllItem(ctx, paging)
}

// GetItemById implements domain.TodoItemUseCase.
func (useCase *todoUsecase) GetItemById(ctx context.Context, conditions map[string]interface{}) (*domain.ToDoItem, error) {
	return useCase.taskRepo.GetItemById(ctx, conditions)
}

// UpdateItem implements domain.TodoItemUseCase.
func (useCase *todoUsecase) UpdateItem(ctx context.Context, conditions map[string]interface{}, data *domain.ToDoItem, status string) error {
	if data.Status == "Finished" {
		logger.LogErrorNoContext("can not update finished item")
		return errors.New("can not update finished item")
	}
	return useCase.taskRepo.UpdateItem(ctx, conditions, data, status)
}
