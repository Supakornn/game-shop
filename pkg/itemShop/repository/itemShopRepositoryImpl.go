package repository

type itemShopRepositoryImpl struct {
}

func NewItemShopRepositoryImpl() ItemShopRepository {
	return &itemShopRepositoryImpl{}
}
