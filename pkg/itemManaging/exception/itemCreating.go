package exception

type ItemCreating struct {
}

func (e *ItemCreating) Error() string {
	return "Creating item failed"
}
