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
	UserID      int       `json:"user_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

func NewSQLiteStore(dbPath string) (*ExpenseStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create tables with user scoping
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			amount REAL NOT NULL,
			description TEXT NOT NULL,
			category TEXT NOT NULL,
			date TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
		CREATE INDEX IF NOT EXISTS idx_expenses_user ON expenses(user_id);
		CREATE INDEX IF NOT EXISTS idx_expenses_date ON expenses(date);
	`)
	if err != nil {
		return nil, err
	}

	// Seed a default user for demo
	db.Exec("INSERT OR IGNORE INTO users (id, username, password_hash) VALUES (1, 'demo', 'demo123')")

	return &ExpenseStore{db: db}, nil
}

func (s *ExpenseStore) CreateUser(username, password string) (User, error) {
	res, err := s.db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, password)
	if err != nil {
		return User{}, err
	}
	id, _ := res.LastInsertId()
	return User{ID: int(id), Username: username}, nil
}

func (s *ExpenseStore) GetUserByCredentials(username, password string) (User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ? AND password_hash = ?", username, password).Scan(&u.ID, &u.Username, &u.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *ExpenseStore) Add(exp Expense) (Expense, error) {
	result, err := s.db.Exec(
		"INSERT INTO expenses (user_id, amount, description, category, date) VALUES (?, ?, ?, ?, ?)",
		exp.UserID, exp.Amount, exp.Description, exp.Category, exp.Date.Format(time.RFC3339),
	)
	if err != nil {
		return Expense{}, err
	}
	id, _ := result.LastInsertId()
	exp.ID = int(id)
	return exp, nil
}

func (s *ExpenseStore) ListByUser(userID int) ([]Expense, error) {
	rows, err := s.db.Query("SELECT id, user_id, amount, description, category, date FROM expenses WHERE user_id = ? ORDER BY date DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		var dateStr string
		if err := rows.Scan(&e.ID, &e.UserID, &e.Amount, &e.Description, &e.Category, &dateStr); err != nil {
			return nil, err
		}
		e.Date, _ = time.Parse(time.RFC3339, dateStr)
		expenses = append(expenses, e)
	}
	return expenses, nil
}

func (s *ExpenseStore) Update(exp Expense) (Expense, error) {
	_, err := s.db.Exec(
		"UPDATE expenses SET amount=?, description=?, category=?, date=? WHERE id=? AND user_id=?",
		exp.Amount, exp.Description, exp.Category, exp.Date.Format(time.RFC3339), exp.ID, exp.UserID,
	)
	if err != nil {
		return Expense{}, err
	}
	return exp, nil
}

func (s *ExpenseStore) Delete(id, userID int) error {
	_, err := s.db.Exec("DELETE FROM expenses WHERE id=? AND user_id=?", id, userID)
	return err
}

func (s *ExpenseStore) SummaryByUser(userID int) (map[string]float64, float64, error) {
	rows, err := s.db.Query("SELECT category, amount FROM expenses WHERE user_id = ?", userID)
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