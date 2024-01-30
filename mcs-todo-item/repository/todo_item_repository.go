package repository

import (
	"context"
	"mcs-nghiadeptrai/mcs-todo-item/domain"

	"gorm.io/gorm"
)

type todoItemRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.ToDoItemRepository {
	return &todoItemRepository{db: db}
}

func (s *todoItemRepository) CreateItem(ctx context.Context, data *domain.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

// DeleteItem implements domain.TodoItemUseCase.
func (s *todoItemRepository) DeleteItem(ctx context.Context, conditions map[string]interface{}) error {
	if err := s.db.
		Table(domain.ToDoItem{}.TableName()).
		Where(conditions).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}

// GetAllItem implements domain.TodoItemUseCase.
func (s *todoItemRepository) GetAllItem(ctx context.Context, paging *domain.DataPaging) ([]domain.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []domain.ToDoItem

	if err := s.db.Table(domain.ToDoItem{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// GetItemById implements domain.TodoItemUseCase.
func (s *todoItemRepository) GetItemById(ctx context.Context, conditions map[string]interface{}) (*domain.ToDoItem, error) {
	var itemData domain.ToDoItem

	if err := s.db.Where(conditions).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, domain.ErrItemNotFound
		}

		return nil, err // other errors
	}

	return &itemData, nil
}

// UpdateItem implements domain.TodoItemUseCase.
func (s *todoItemRepository) UpdateItem(ctx context.Context, conditions map[string]interface{}, data *domain.ToDoItem, status string) error {
	if err := s.db.Where(conditions).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
