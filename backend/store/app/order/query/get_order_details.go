package query

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

func (*query) GetOrderDetails(id string) (*Order, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	record, err := queryFactory.
		New(&data.Order{}).
		Include("Customer").
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(record.(*data.Order))

	return &order, nil
}
