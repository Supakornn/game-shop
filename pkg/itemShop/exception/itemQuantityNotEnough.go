package exception

import "fmt"

type ItemQuantityNotEnough struct {
	ItemID uint64
}

func (e *ItemQuantityNotEnough) Error() string {
	return fmt.Sprintf("item quantity not enough: %d", e.ItemID)
}
