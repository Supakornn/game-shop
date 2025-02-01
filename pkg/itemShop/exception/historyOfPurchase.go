package exception

type HistoryOfPurchase struct{}

func (h *HistoryOfPurchase) Error() string {
	return "history of purchase not found"
}
