package postgres

import "store/app/data"

type stockRepository struct {
	postgresRepository
}

func (r *stockRepository) GetByOrderItems(
	items []data.OrderItem,
	tx data.Transaction,
) (data.Stock, error) {
	bookIds := []data.EntityId{}
	for _, item := range items {
		bookIds = append(bookIds, item.BookId)
	}

	records, err := r.
		Query(&data.Book{}, tx).
		Select("id", "onhand_qty", "preserved_qty").
		Where("id IN ?", bookIds).
		Find()

	if err != nil {
		return nil, err
	}

	stock := data.Stock{}
	for _, record := range records {
		book := record.(data.Book)
		stock[book.Id] = 
	}
}
