package todostorage

import (
	"context"
	todomodel "mcs-nghiadeptrai/mcs-user/module/item/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
