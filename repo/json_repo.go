package repo

import "github.com/Brownie44l1/cli_inventory/model"

type InventoryRepo interface {
	SaveItem(item model.Item) error
	GetAllItems() ([]model.Item, error)
	FindBySKU(sku string) (*model.Item, error)
}