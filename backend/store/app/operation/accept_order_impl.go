package operation

import (
	"store/app/data"
	tx "store/app/domain/transaction"
)

type acceptOrder struct{}

func (*acceptOrder) Accept(id string) error {
	if err := tx.AcceptOrder(data.FromStringToEntityId(id)); err != nil {
		return err
	}

	return nil
}
