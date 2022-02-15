package query

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

func (*query) GetOrderDetails(id string) (*Order, error) {
	var queryFactory = container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	orderId := data.FromStringToEntityId(id)
	record, err := queryFactory.
		New(&data.Order{}).
		Include("Customer").
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id").Eq(orderId).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(record.(*data.Order))

	return &order, nil
}
