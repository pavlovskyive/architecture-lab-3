package orders

import (
	"database/sql"
	"fmt"
)

// Order stores info about one order
type Order struct {
	ID          int64 `json:"id"`
	MenuItemID  int64 `json:"menu_item_id"`
	TableNumber uint  `json:"table_number"`
}

// Store handles pointer to DB
type Store struct {
	Db *sql.DB
}

// NewStore is a function that forms Store structure from db
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// ListOrders returns array of all items in orders table of DB
func (s *Store) ListOrders() ([]*Order, error) {
	rows, err := s.Db.Query("SELECT id, menu_item_id, table_number FROM orders LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Order

	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.MenuItemID, &o.TableNumber); err != nil {
			return nil, err
		}
		res = append(res, &o)
	}

	if res == nil {
		res = make([]*Order, 0)
	}

	return res, nil
}

// CreateOrder inserts a menu item to DB
func (s *Store) CreateOrder(menuItemID int64, tableNumber uint) error {
	if menuItemID == 0 {
		return fmt.Errorf("menu item id is not provided")
	}

	if tableNumber == 0 {
		return fmt.Errorf("table number is not provided")
	}

	_, err := s.Db.Exec("INSERT INTO orders (menu_item_id, table_number) VALUES ($1, $2)", menuItemID, tableNumber)
	return err
}
