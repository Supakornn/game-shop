package exception

type CoinNotEnough struct{}

func (e *CoinNotEnough) Error() string {
	return "coin not enough"
}
