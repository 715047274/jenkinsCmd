package domain

type Inventory struct {
	ProductID int64
	Stock     int
}

type InventoryEvent struct {
	ProductID int64
	Quantity  int
}

func (e InventoryEvent) EventType() string {
	return "InventoryUpdated"
}
