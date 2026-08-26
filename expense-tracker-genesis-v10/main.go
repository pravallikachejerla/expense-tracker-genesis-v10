package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"expense-tracker/internal/service"
	"expense-tracker/internal/store"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	dbPath := "expenses.db"
	st, err := store.NewSQLiteStore(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	svc := service.NewService(st)

	// Seed sample data if table is empty
	expenses, _ := svc.ListExpenses()
	if len(expenses) == 0 {
		svc.AddExpense(12.5, "Lunch", "Food", time.Now().Add(-24*time.Hour))
		svc.AddExpense(45.0, "Taxi", "Transport", time.Now().Add(-2*time.Hour))
		svc.AddExpense(1200.0, "August Rent", "Rent", time.Now().Add(-30*24*time.Hour))
		log.Println("Sample data seeded")
	}

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/expenses", func(w http.ResponseWriter, r *http.Request) {
		handleExpenses(w, r, svc)
	})
	mux.HandleFunc("/api/expenses/", func(w http.ResponseWriter, r *http.Request) {
		handleExpenseByID(w, r, svc)
	})
	mux.HandleFunc("/api/summary", func(w http.ResponseWriter, r *http.Request) {
		handleSummary(w, r, svc)
	})
	mux.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		handleCategories(w, r, svc)
	})

	// Serve frontend
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "frontend/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Expense Tracker running on http://localhost:%s", port)
	log.Println("Open browser to http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func jsonResponse(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func handleExpenses(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	switch r.Method {
	case http.MethodGet:
		category := r.URL.Query().Get("category")
		startStr := r.URL.Query().Get("start")
		endStr := r.URL.Query().Get("end")

		var expenses []service.Expense
		var err error

		if category != "" {
			expenses, err = svc.ListByCategory(category)
		} else if startStr != "" && endStr != "" {
			start, _ := time.Parse(time.RFC3339, startStr)
			end, _ := time.Parse(time.RFC3339, endStr)
			expenses, err = svc.ListByDateRange(start, end)
		} else {
			expenses, err = svc.ListExpenses()
		}

		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusOK, APIResponse{Success: true, Data: expenses})

	case http.MethodPost:
		var req struct {
			Amount      float64   `json:"amount"`
			Description string    `json:"description"`
			Category    string    `json:"category"`
			Date        time.Time `json:"date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}

		exp, err := svc.AddExpense(req.Amount, req.Description, req.Category, req.Date)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusCreated, APIResponse{Success: true, Data: exp})

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, APIResponse{Success: false, Error: "method not allowed"})
	}
}

func handleExpenseByID(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/expenses/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: "invalid id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		exp, err := svc.GetExpense(id)
		if err != nil {
			jsonResponse(w, http.StatusNotFound, APIResponse{Success: false, Error: "not found"})
			return
		}
		jsonResponse(w, http.StatusOK, APIResponse{Success: true, Data: exp})

	case http.MethodPut:
		var req struct {
			Amount      float64   `json:"amount"`
			Description string    `json:"description"`
			Category    string    `json:"category"`
			Date        time.Time `json:"date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}

		exp, err := svc.UpdateExpense(id, req.Amount, req.Description, req.Category, req.Date)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusOK, APIResponse{Success: true, Data: exp})

	case http.MethodDelete:
		if err := svc.DeleteExpense(id); err != nil {
			jsonResponse(w, http.StatusInternalServerError, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusOK, APIResponse{Success: true})

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, APIResponse{Success: false, Error: "method not allowed"})
	}
}

func handleSummary(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusMethodNotAllowed, APIResponse{Success: false, Error: "method not allowed"})
		return
	}

	byCat, total, err := svc.GetSummary()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, APIResponse{Success: false, Error: err.Error()})
		return
	}
	jsonResponse(w, http.StatusOK, APIResponse{Success: true, Data: map[string]interface{}{
		"total":       total,
		"by_category": byCat,
	}})
}

func handleCategories(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	switch r.Method {
	case http.MethodGet:
		cats, err := svc.GetCategories()
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusOK, APIResponse{Success: true, Data: cats})

	case http.MethodPost:
		var req struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}
		if err := svc.AddCategory(req.Name); err != nil {
			jsonResponse(w, http.StatusBadRequest, APIResponse{Success: false, Error: err.Error()})
			return
		}
		jsonResponse(w, http.StatusCreated, APIResponse{Success: true})

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, APIResponse{Success: false, Error: "method not allowed"})
	}
}
