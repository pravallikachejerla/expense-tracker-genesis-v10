package service

import (
	"fmt"
	"time"

	"expense-tracker/internal/store"
)

type Service struct {
	store *store.ExpenseStore
}

func NewService(s *store.ExpenseStore) *Service {
	return &Service{store: s}
}

type Expense = store.Expense

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (s *Service) Register(username, password string) (User, error) {
	if username == "" || password == "" {
		return User{}, fmt.Errorf("username and password required")
	}
	// In full impl: hash password, check uniqueness
	u, err := s.store.CreateUser(username, password)
	if err != nil {
		return User{}, err
	}
	return User{ID: u.ID, Username: u.Username}, nil
}

func (s *Service) Login(username, password string) (string, error) {
	user, err := s.store.GetUserByCredentials(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	// Simplified JWT (no real signing in this stub; real would use claims)
	token := fmt.Sprintf("fake-jwt-for-user-%d", user.ID) // Replace with real jwt.New in production
	return token, nil
}

func (s *Service) AddExpense(amount float64, desc, cat string, date time.Time) (Expense, error) {
	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be positive")
	}
	if date.IsZero() {
		date = time.Now()
	}
	return s.store.Add(store.Expense{
		Amount:      amount,
		Description: desc,
		Category:    cat,
		Date:        date,
	})
}

func (s *Service) ListExpenses(userID int) ([]Expense, error) {
	return s.store.ListByUser(userID)
}

func (s *Service) UpdateExpense(e Expense) (Expense, error) {
	return s.store.Update(e)
}

func (s *Service) DeleteExpense(id, userID int) error {
	return s.store.Delete(id, userID)
}

func (s *Service) GetSummary(userID int) (map[string]float64, float64, error) {
	return s.store.SummaryByUser(userID)
}