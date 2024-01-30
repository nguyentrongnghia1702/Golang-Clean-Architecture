package todostorage

import (
	"context"
	todomodel "mcs-nghiadeptrai/mcs-user/module/item/model"
)

func (s *mysqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {

	if err := s.db.
		Table(todomodel.ToDoItem{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
