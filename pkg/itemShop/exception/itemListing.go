package exception

type ItemListing struct{}

func (e *ItemListing) Error() string {
	return "Item listing error"
}
