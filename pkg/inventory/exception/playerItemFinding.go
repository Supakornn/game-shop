package exception

import "fmt"

type PlayerItemFinding struct {
	PlayerID string
}

func (e *PlayerItemFinding) Error() string {
	return fmt.Sprintf("player item finding failed: %s", e.PlayerID)
}
