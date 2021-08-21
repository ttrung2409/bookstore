package postgres

import (
	"errors"
	"fmt"
	"reflect"
	data "store/app/domain/data"
	repo "store/app/repository"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct {
	newEntity func() data.Entity
}

func (r *postgresRepository) Query(model interface{}, tx repo.Transaction) repo.Query {
	return newQuery(model, tx)
}

func (r *postgresRepository) get(id data.EntityId, tx repo.Transaction) (interface{}, error) {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	entity := r.newEntity()
	key := getPrimaryKey(entity)
	if key == "" {
		return nil, errors.New("No primary key found")
	}

	if result := db.Where(fmt.Sprintf("%s = ?", key)).Find(entity); result.Error != nil {
		return nil, toDataQueryError(result.Error)
	}

	return entity, nil
}

func (r *postgresRepository) create(
	entity data.Entity,
	tx repo.Transaction,
) (data.EntityId, error) {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Omit(clause.Associations).Create(entity); result.Error != nil {
		return data.EmptyEntityId, result.Error
	}

	return entity.GetId(), nil
}

func (r *postgresRepository) update(
	id data.EntityId,
	entity data.Entity,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	entity.SetId(id)
	result := db.Model(entity).Omit(clause.Associations).Updates(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getPrimaryKey(entity data.Entity) string {
	entityType := reflect.TypeOf(entity).Elem()
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		if strings.Contains(field.Tag.Get("gorm"), "primaryKey") {
			return field.Name
		}
	}

	return ""
}

func toDataQueryError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repo.ErrNotFound
	}

	return err
}
