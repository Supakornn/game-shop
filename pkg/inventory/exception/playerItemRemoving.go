package exception

import "fmt"

type PlayerItemRemoving struct {
	ItemID uint64
}

func (e *PlayerItemRemoving) Error() string {
	return fmt.Sprintf("player item removing failed: %d", e.ItemID)
}
