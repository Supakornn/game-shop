package exception

import "fmt"

type InventoryFilling struct {
	PlayerID string
	ItemID   uint64
}

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("filling inventory failed: playerID: %s, itemID: %d", e.PlayerID, e.ItemID)
}
