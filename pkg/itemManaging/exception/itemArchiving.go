package exception

import "fmt"

type ItemArchiving struct {
	ItemID uint64
}

func (e *ItemArchiving) Error() string {
	return fmt.Sprintf("Archiving item with ID %d failed", e.ItemID)
}
