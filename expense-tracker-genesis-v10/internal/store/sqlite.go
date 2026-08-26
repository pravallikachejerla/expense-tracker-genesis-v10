package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type ExpenseStore struct {
	db *sql.DB
}

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

func NewSQLiteStore(dbPath string) (*ExpenseStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			amount REAL NOT NULL,
			description TEXT NOT NULL,
			category TEXT NOT NULL,
			date TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			name TEXT PRIMARY KEY
		)
	`)
	if err != nil {
		return nil, err
	}

	// Seed default categories
	defaultCats := []string{"Food", "Transport", "Rent", "Entertainment", "Utilities", "Other"}
	for _, c := range defaultCats {
		db.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", c)
	}

	return &ExpenseStore{db: db}, nil
}

func (s *ExpenseStore) Add(exp Expense) (Expense, error) {
	result, err := s.db.Exec(
		"INSERT INTO expenses (amount, description, category, date) VALUES (?, ?, ?, ?)",
		exp.Amount, exp.Description, exp.Category, exp.Date.Format(time.RFC3339),
	)
	if err != nil {
		return Expense{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Expense{}, err
	}

	exp.ID = int(id)
	return exp, nil
}

func (s *ExpenseStore) GetByID(id int) (Expense, error) {
	var e Expense
	var dateStr string
	err := s.db.QueryRow("SELECT id, amount, description, category, date FROM expenses WHERE id = ?", id).
		Scan(&e.ID, &e.Amount, &e.Description, &e.Category, &dateStr)
	if err != nil {
		return Expense{}, err
	}
	e.Date, _ = time.Parse(time.RFC3339, dateStr)
	return e, nil
}

func (s *ExpenseStore) List() ([]Expense, error) {
	rows, err := s.db.Query("SELECT id, amount, description, category, date FROM expenses ORDER BY date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		var dateStr string
		if err := rows.Scan(&e.ID, &e.Amount, &e.Description, &e.Category, &dateStr); err != nil {
			return nil, err
		}
		e.Date, _ = time.Parse(time.RFC3339, dateStr)
		expenses = append(expenses, e)
	}
	return expenses, nil
}

func (s *ExpenseStore) ListByCategory(category string) ([]Expense, error) {
	rows, err := s.db.Query("SELECT id, amount, description, category, date FROM expenses WHERE category = ? ORDER BY date DESC", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		var dateStr string
		if err := rows.Scan(&e.ID, &e.Amount, &e.Description, &e.Category, &dateStr); err != nil {
			return nil, err
		}
		e.Date, _ = time.Parse(time.RFC3339, dateStr)
		expenses = append(expenses, e)
	}
	return expenses, nil
}

func (s *ExpenseStore) ListByDateRange(start, end time.Time) ([]Expense, error) {
	rows, err := s.db.Query(`
		SELECT id, amount, description, category, date 
		FROM expenses 
		WHERE date >= ? AND date <= ? 
		ORDER BY date DESC`, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		var dateStr string
		if err := rows.Scan(&e.ID, &e.Amount, &e.Description, &e.Category, &dateStr); err != nil {
			return nil, err
		}
		e.Date, _ = time.Parse(time.RFC3339, dateStr)
		expenses = append(expenses, e)
	}
	return expenses, nil
}

func (s *ExpenseStore) Update(exp Expense) (Expense, error) {
	_, err := s.db.Exec(
		"UPDATE expenses SET amount=?, description=?, category=?, date=? WHERE id=?",
		exp.Amount, exp.Description, exp.Category, exp.Date.Format(time.RFC3339), exp.ID,
	)
	if err != nil {
		return Expense{}, err
	}
	return exp, nil
}

func (s *ExpenseStore) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM expenses WHERE id = ?", id)
	return err
}

func (s *ExpenseStore) Summary() (map[string]float64, float64, error) {
	rows, err := s.db.Query("SELECT category, amount FROM expenses")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	total := 0.0
	byCat := make(map[string]float64)
	for rows.Next() {
		var cat string
		var amt float64
		if err := rows.Scan(&cat, &amt); err != nil {
			return nil, 0, err
		}
		total += amt
		byCat[cat] += amt
	}
	return byCat, total, nil
}

func (s *ExpenseStore) GetCategories() ([]string, error) {
	rows, err := s.db.Query("SELECT name FROM categories ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		cats = append(cats, name)
	}
	return cats, nil
}

func (s *ExpenseStore) AddCategory(name string) error {
	_, err := s.db.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", name)
	return err
}

func (s *ExpenseStore) Close() error {
	return s.db.Close()
}
