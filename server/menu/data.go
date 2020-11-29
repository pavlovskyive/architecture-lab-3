package menu

import (
	"database/sql"
	"fmt"
)

// Item stores info about one item from menu
type Item struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

// Store handles pointer to DB
type Store struct {
	Db *sql.DB
}

// NewStore is a function that forms Store structure from db
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// ListMenuItems returns array of all items in menu table of DB
func (s *Store) ListMenuItems() ([]*Item, error) {
	rows, err := s.Db.Query("SELECT id, name, price FROM menu LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Item

	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		res = append(res, &i)
	}

	if res == nil {
		res = make([]*Item, 0)
	}

	return res, nil
}

// CreateMenuItem inserts a menu item to DB
func (s *Store) CreateMenuItem(name string, price uint) error {
	if len(name) == 0 {
		return fmt.Errorf("menu item name is not provided")
	}

	if price == 0 {
		return fmt.Errorf("menu item price is not provided")
	}

	_, err := s.Db.Exec("INSERT INTO menu (name, price) VALUES ($1, $2)", name, price)
	return err
}
