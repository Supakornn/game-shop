package exception

type CoinAdding struct{}

func (e *CoinAdding) Error() string {
	return "Error adding coin"
}
