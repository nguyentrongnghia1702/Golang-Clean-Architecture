package todobiz

import (
	"context"
	"errors"
	"mcs-nghiadeptrai/mcs-common/logger"
	todomodel "mcs-nghiadeptrai/mcs-user/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if data.Title == "" {
		logger.LogErrorNoContext("title can not be blank")
		return errors.New("title can not be blank")
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		logger.LogErrorNoContext("create new item failed")
		return err
	}
	return nil
}
