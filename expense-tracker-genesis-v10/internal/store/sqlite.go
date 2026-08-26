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

func (s *ExpenseStore) List() ([]Expense, error) {
	rows, err := s.db.Query("SELECT id, amount, description, category, date FROM expenses")
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

func (s *ExpenseStore) Close() error {
	return s.db.Close()
}
