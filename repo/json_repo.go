package repo

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Brownie44l1/cli_inventory/model"
)

type InventoryRepo interface {
	SaveItem(item model.Item) error
	GetAllItems() ([]model.Item, error)
	FindBySKU(sku string) (*model.Item, error)
}

type FileStore struct {
	//inventoryRepo InventoryRepo
	filepath      string
}

func NewFileStore(filepath string) *FileStore {
	return &FileStore{
		//inventoryRepo: inventoryRepo,
		filepath:      filepath,
	}
}

func (fs *FileStore) GetAllItems() ([]model.Item, error) {
	var items []model.Item

	data, err := os.ReadFile(fs.filepath)
	if os.IsNotExist(err) {
		return []model.Item{}, nil
	}
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		if err = json.Unmarshal(data, &items); err != nil {
			return nil, err
		}
	}
	return items, nil
}

func (fs *FileStore) SaveItem(item model.Item) error {
	items, err := fs.GetAllItems()
	if err != nil {
		return err
	}

	maxID := 0
	for _, it := range items {
		if it.ID > maxID {
			maxID = it.ID
		}
	}

	item.ID = maxID + 1
	items = append(items, item)

	data, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(fs.filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) FindBySKU(sku string) (*model.Item, error) {
	items, err := fs.GetAllItems()
	if err != nil {
		return nil, err
	}

	for i := range items {
		if items[i].Sku == sku {
			return &items[i], nil
		}
	}
	return nil, errors.New("SKU not found")
}
