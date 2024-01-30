package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string { return "ToDoItems" }

type ToDoItemRepository interface {
	CreateItem(ctx context.Context, data *ToDoItem) error
	DeleteItem(ctx context.Context, conditions map[string]interface{}) error
	UpdateItem(ctx context.Context, conditions map[string]interface{}, data *ToDoItem, status string) error
	GetAllItem(ctx context.Context, paging *DataPaging) ([]ToDoItem, error)
	GetItemById(ctx context.Context, conditions map[string]interface{}) (*ToDoItem, error)
}

type TodoItemUseCase interface {
	CreateItem(ctx context.Context, data *ToDoItem) error
	DeleteItem(ctx context.Context, conditions map[string]interface{}) error
	UpdateItem(ctx context.Context, conditions map[string]interface{}, data *ToDoItem, status string) error
	GetAllItem(ctx context.Context, paging *DataPaging) ([]ToDoItem, error)
	GetItemById(ctx context.Context, conditions map[string]interface{}) (*ToDoItem, error)
}
