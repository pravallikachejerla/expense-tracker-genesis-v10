package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

type Summary struct {
	Total      float64            `json:"total"`
	ByCategory map[string]float64 `json:"by_category"`
}

var (
	expenses  = make([]Expense, 0)
	nextID    = 1
	mu        sync.Mutex
	categories = []string{"Food", "Transport", "Rent", "Entertainment", "Utilities", "Other"}
)

func isValidCategory(cat string) bool {
	for _, c := range categories {
		if c == cat {
			return true
		}
	}
	return false
}

func addExpense(amount float64, desc, cat string, date time.Time) (Expense, error) {
	mu.Lock()
	defer mu.Unlock()

	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be positive")
	}
	if !isValidCategory(cat) {
		return Expense{}, fmt.Errorf("invalid category: %s. Allowed: %v", cat, categories)
	}
	if date.IsZero() {
		date = time.Now()
	}

	e := Expense{
		ID:          nextID,
		Amount:      amount,
		Description: desc,
		Category:    cat,
		Date:        date,
	}
	expenses = append(expenses, e)
	nextID++
	return e, nil
}

func listExpenses() []Expense {
	mu.Lock()
	defer mu.Unlock()
	return append([]Expense{}, expenses...)
}

func getSummary() Summary {
	mu.Lock()
	defer mu.Unlock()

	total := 0.0
	byCat := make(map[string]float64)
	for _, e := range expenses {
		total += e.Amount
		byCat[e.Category] += e.Amount
	}
	return Summary{Total: total, ByCategory: byCat}
}

func expensesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var e struct {
			Amount      float64   `json:"amount"`
			Description string    `json:"description"`
			Category    string    `json:"category"`
			Date        time.Time `json:"date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		expense, err := addExpense(e.Amount, e.Description, e.Category, e.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expense)
		return
	}

	if r.Method == http.MethodGet {
		list := listExpenses()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	s := getSummary()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/expenses", expensesHandler)
	http.HandleFunc("/summary", summaryHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Seed some sample data
	addExpense(12.5, "Lunch", "Food", time.Now().Add(-24*time.Hour))
	addExpense(45.0, "Taxi", "Transport", time.Now().Add(-2*time.Hour))
	addExpense(1200.0, "August Rent", "Rent", time.Now().Add(-30*24*time.Hour))

	log.Printf("Expense Tracker running on :%s", port)
	log.Println("Endpoints: /expenses (GET/POST), /summary (GET)")
	log.Println("Sample data seeded. Try curl commands.")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
