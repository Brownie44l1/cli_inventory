package service

import (
	"errors"
	"fmt"

	"github.com/Brownie44l1/cli_inventory/model"
	"github.com/Brownie44l1/cli_inventory/repo"
)

type Inventory interface {
	AddItem(item model.Item) error
	ListItems() ([]model.Item, error)
}

type InventoryService struct {
	repo repo.InventoryRepo
}

func NewInventoryService(r repo.InventoryRepo) *InventoryService {
	return &InventoryService{repo: r}
}

func (s *InventoryService) AddItem(item model.Item) error {
	if item.Sku == "" || item.Name == "" {
		return errors.New("missing required field")
	}
	
	if item.Quantity < 0 {
		return errors.New("invalid quantity")
	}
	
	found, err := s.repo.FindBySKU(item.Sku)
    if err == nil && found != nil {
        return fmt.Errorf("an item with SKU %s already exists", item.Sku)
    }
    if err != nil && err.Error() != "SKU not found" {
        return err
    }
	return s.repo.SaveItem(item)
}
