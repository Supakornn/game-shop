package exception

type ItemCounting struct {
}

func (e *ItemCounting) Error() string {
	return "Failed to count item"
}
