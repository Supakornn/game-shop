package exception

import "fmt"

type ItemEditing struct {
	ItemID uint64
}

func (e *ItemEditing) Error() string {
	return fmt.Sprintf("Item with ID %d is being edited", e.ItemID)
}
