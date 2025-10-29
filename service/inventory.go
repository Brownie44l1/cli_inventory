package service

import "github.com/Brownie44l1/cli_inventory/model"

type Inventory interface {
	AddItem(item model.Item) error
	ListItems()([]model.Item, error)
}